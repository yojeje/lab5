package main


import (
	"context"
	"fmt"
	"log"
	"net"
	"sync"

	pb "github.com/yojeje/lab5"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedDataNodeServer
	mu         sync.Mutex
	decisiones map[int32]map[int32]int32 // [id][piso]decision
}

func (s *server) GuardarDecision(ctx context.Context, req *pb.GuardarDecisionRequest) (*pb.GuardarDecisionResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.decisiones == nil {
		s.decisiones = make(map[int32]map[int32]int32)
	}

	if s.decisiones[req.Id] == nil {
		s.decisiones[req.Id] = make(map[int32]int32)
	}

	s.decisiones[req.Id][req.Piso] = req.Decision
	fmt.Printf("Decision guardada: id=%d, piso=%d, decision=%d\n", req.Id, req.Piso, req.Decision)

	return &pb.GuardarDecisionResponse{Success: true}, nil
}

func (s *server) ObtenerDecision(ctx context.Context, req *pb.ObtenerDecisionRequest) (*pb.ObtenerDecisionResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.decisiones == nil || s.decisiones[req.Id] == nil {
		return &pb.ObtenerDecisionResponse{Decision: -1}, fmt.Errorf("no hay decision para id %d y piso %d", req.Id, req.Piso)
	}

	decision, ok := s.decisiones[req.Id][req.Piso]
	if !ok {
		return &pb.ObtenerDecisionResponse{Decision: -1}, fmt.Errorf("no hay decision para id %d y piso %d", req.Id, req.Piso)
	}

	return &pb.ObtenerDecisionResponse{Decision: decision}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50054")
	if err != nil {
		log.Fatalf("Error al escuchar: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterDataNodeServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Error al servir: %v", err)
	}
}
