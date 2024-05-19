package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"net"
	"sync"
	"time"

	pb "github.com/yojeje/lab5"

	"github.com/streadway/amqp"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct {
	pb.UnimplementedDirectorServer
	mu             sync.Mutex
	mercenarios    map[int32]bool
	montoAcumulado int64
	rng            *rand.Rand
	rabbitConn     *amqp.Connection
}

func (s *server) Preparacion(ctx context.Context, in *pb.PreparacionRequest) (*pb.PreparacionResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.mercenarios == nil {
		s.mercenarios = make(map[int32]bool)
	}
	s.mercenarios[in.Id] = true
	log.Printf("Mercenario %d preparado para la mision", in.Id)
	return &pb.PreparacionResponse{Success: true}, nil
}

func (s *server) EnviarDecision(ctx context.Context, in *pb.DecisionRequest) (*pb.DecisionResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	log.Printf("Decisión recibida para Mercenario %d en Piso %d: %d", in.Id, in.Piso, in.Decision)
	resultado := "Sobreviviste"
	switch in.Piso {
	case 1:
		resultado = s.calcularProbabilidadesPiso1(in.Decision)
	case 2:
		resultado = s.seleccionarPasilloPiso2(in.Decision)
	case 3:
		resultado = s.confrontacionFinalPiso3(in.Decision)
	}

	if resultado == "Eliminado" {
		delete(s.mercenarios, in.Id)
		s.montoAcumulado += 100000000
		log.Printf("Mercenario %d ha sido eliminado", in.Id)
		if err := s.informarEliminacion(in.Id); err != nil {
			log.Printf("Error al informar eliminación: %v", err)
		}
	}
	return &pb.DecisionResponse{DataNodeIp: resultado}, nil
}

func (s *server) calcularProbabilidadesPiso1(decision int32) string {
	x := s.rng.Intn(100)
	y := s.rng.Intn(100)
	for y == x {
		y = s.rng.Intn(100)
	}
	if x > y {
		x, y = y, x
	}

	var probabilidad int
	switch decision {
	case 1:
		probabilidad = x
	case 2:
		probabilidad = y - x
	case 3:
		probabilidad = 100 - y
	}

	if s.rng.Intn(100) < probabilidad {
		return "Sobreviviste"
	}
	return "Eliminado"
}

func (s *server) seleccionarPasilloPiso2(decision int32) string {
	correctPath := s.rng.Intn(2) + 1
	if decision == int32(correctPath) {
		return "Sobreviviste"
	}
	return "Eliminado"
}

func (s *server) confrontacionFinalPiso3(decision int32) string {
	aciertos := 0
	for i := 0; i < 5; i++ {
		if decision == int32(s.rng.Intn(15)+1) {
			aciertos++
		}
	}
	if aciertos >= 2 {
		return "Sobreviviste"
	}
	return "Eliminado"
}

func (s *server) informarEliminacion(mercenarioId int32) error {
	if s.rabbitConn == nil {
		return errors.New("no se ha establecido la conexión con RabbitMQ")
	}

	ch, err := s.rabbitConn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"dosh_bank_exchange",
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	mensaje := fmt.Sprintf("Mercenario eliminado: %d", mercenarioId)

	err = ch.Publish(
		"dosh_bank_exchange",
		"eliminacion",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(mensaje),
		},
	)
	if err != nil {
		return err
	}

	return nil
}

func (s *server) connectRabbitMQ() error {
	conn, err := amqp.Dial("amqp://guest:guest@localhost")
	if err != nil {
		return err
	}
	s.rabbitConn = conn
	return nil
}

func main() {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	directorServer := &server{rng: rng}
	if err := directorServer.connectRabbitMQ(); err != nil {
		log.Fatalf("failed to connect to RabbitMQ: %v", err)
	}
	defer directorServer.rabbitConn.Close()

	pb.RegisterDirectorServer(s, &server{})
	log.Printf("Director server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
