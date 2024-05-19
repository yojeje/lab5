package main

import (
	"context"
	"log"
	"net"
	"sync"

	pb "github.com/yojeje/lab5"

	"google.golang.org/grpc"
)

const (
	port = ":50000"
)

type server struct {
	mu                sync.Mutex
	decisionesPiso1   map[int32]int32 // Mapa para almacenar las decisiones del piso 1
	decisionesPiso2   map[int32]int32 // Mapa para almacenar las decisiones del piso 2
	decisionesPiso3   map[int32]int32 // Mapa para almacenar las decisiones del piso 3
	dataNodeAddresses []string        // Direcciones de los DataNode
	pb.UnimplementedNameNodeServer
}

func (s *server) RegistrarDecision(ctx context.Context, req *pb.DecisionRequest) (*pb.DecisionResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Registrar la decisión del mercenario para el piso correspondiente
	switch req.Piso {
	case 1:
		s.decisionesPiso1[req.Id] = req.Decision
	case 2:
		s.decisionesPiso2[req.Id] = req.Decision
	case 3:
		s.decisionesPiso3[req.Id] = req.Decision
	default:
		log.Printf("Piso no válido: %d", req.Piso)
		return nil, nil
	}

	// Distribuir la decisión entre los DataNode
	dataNodeIndex := int(req.Id) % len(s.dataNodeAddresses)
	dataNodeAddress := s.dataNodeAddresses[dataNodeIndex]

	// Conectar y enviar la decisión al DataNode correspondiente usando gRPC
	conn, err := grpc.Dial(dataNodeAddress, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("No se pudo conectar al DataNode: %v", err)
		return nil, err
	}
	defer conn.Close()

	dataNodeClient := pb.NewDataNodeClient(conn)
	_, err = dataNodeClient.GuardarDecision(ctx, &pb.GuardarDecisionRequest{
	Id:       req.Id,
	Piso:     req.Piso,
	Decision: req.Decision,
	})

	if err != nil {
		log.Printf("Error al guardar decisión en DataNode: %v", err)
		return nil, err
	}

	return &pb.DecisionResponse{DataNodeIp: dataNodeAddress}, nil
}

func (s *server) ObtenerDecision(ctx context.Context, req *pb.ObtenerDecisionRequest) (*pb.ObtenerDecisionResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Obtener la decisión del mercenario para el piso correspondiente
	var decision int32
	switch req.Piso {
	case 1:
		decision = s.decisionesPiso1[req.Id]
	case 2:
		decision = s.decisionesPiso2[req.Id]
	case 3:
		decision = s.decisionesPiso3[req.Id]
	default:
		log.Printf("Piso no válido: %d", req.Piso)
		return nil, nil
	}

	return &pb.ObtenerDecisionResponse{Decision: decision}, nil
}

func main() {
	// Direcciones de los DataNode
	dataNodeAddresses := []string{"10.0.0.1", "10.0.0.2", "10.0.0.3"}

	// Inicializar el servidor gRPC
	s := grpc.NewServer()

	// Crear una instancia del servidor NameNode
	server := &server{
		decisionesPiso1:   make(map[int32]int32),
		decisionesPiso2:   make(map[int32]int32),
		decisionesPiso3:   make(map[int32]int32),
		dataNodeAddresses: dataNodeAddresses,
	}

	// Registrar el servidor NameNode en el servidor gRPC
	pb.RegisterNameNodeServer(s, server)

	// Escuchar en el puerto
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("NameNode listening on port %s", port)

	// Servir peticiones gRPC
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
