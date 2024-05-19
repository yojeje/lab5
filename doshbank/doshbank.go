package main

import (
	"context"
	"log"
	"net"
	"os"
	"strconv"
	"sync"

	pb "github.com/yojeje/lab5"
	"github.com/streadway/amqp"
	"google.golang.org/grpc"
)

const (
	port         = ":50053"
	montoFile    = "monto.txt"
	logFile      = "eliminaciones.txt"
	initialMonto = 0

	rabbitMQURL  = "amqp://guest:guest@localhost:5672/"
	exchangeName = "eliminaciones_exchange"
	routingKey   = "eliminaciones"
)

type server struct {
	mu          sync.Mutex
	montoActual int64
	pb.UnimplementedDoshBankServer
}

func (s *server) RegistrarEliminacion(ctx context.Context, req *pb.RegistroRequest) (*pb.RegistroResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Actualizar el archivo de registro de eliminaciones
	file, err := os.OpenFile(logFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	if _, err := file.WriteString("Mercenario ID: " + strconv.Itoa(int(req.Id)) + ", Piso: " + strconv.Itoa(int(req.Piso)) + "\n"); err != nil {
		return nil, err
	}

	// Incrementar el monto acumulado
	s.montoActual += 100000000 // Aumentar en 100 millones de libras
	if err := s.actualizarMontoArchivo(); err != nil {
		return nil, err
	}

	// Enviar mensaje a RabbitMQ en una goroutine
	go enviarMensajeRabbitMQ()

	return &pb.RegistroResponse{}, nil
}

func (s *server) ObtenerMonto(ctx context.Context, req *pb.MontoRequest) (*pb.MontoResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	return &pb.MontoResponse{Monto: s.montoActual}, nil
}

func (s *server) actualizarMontoArchivo() error {
	// Actualizar el archivo de monto acumulado
	if err := os.WriteFile(montoFile, []byte(strconv.FormatInt(s.montoActual, 10)), 0644); err != nil {
		return err
	}
	return nil
}

func enviarMensajeRabbitMQ() {
	// Conexión a RabbitMQ
	conn, err := amqp.Dial(rabbitMQURL)
	if err != nil {
		log.Printf("failed to connect to RabbitMQ: %v", err)
		return
	}
	defer conn.Close()

	// Canal de comunicación
	ch, err := conn.Channel()
	if err != nil {
		log.Printf("failed to open a channel: %v", err)
		return
	}
	defer ch.Close()

	// Declarar intercambio
	err = ch.ExchangeDeclare(
		exchangeName, // nombre
		"fanout",     // tipo
		true,         // durable
		false,        // auto-eliminación
		false,        // interno
		false,        // no esperar
		nil,          // argumentos
	)
	if err != nil {
		log.Printf("failed to declare an exchange: %v", err)
		return
	}

	// Mensaje a enviar
	msg := "Mercenario eliminado"

	// Publicar mensaje
	err = ch.Publish(
		exchangeName, // intercambio
		routingKey,   // clave de enrutamiento
		false,        // obligatorio
		false,        // inmediato
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(msg),
		})
	if err != nil {
		log.Printf("failed to publish a message: %v", err)
		return
	}
}

func main() {
	// Inicializar el servidor gRPC
	s := grpc.NewServer()
	pb.RegisterDoshBankServer(s, &server{montoActual: initialMonto})

	// Escuchar en el puerto
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("Dosh Bank listening on port %s", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
