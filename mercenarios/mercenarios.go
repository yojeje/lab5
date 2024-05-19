package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	pb "github.com/yojeje/lab5"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	address        = "director:50051"
	numMercenarios = 8
)

var mercenarioID int = 8

func main() {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewDirectorClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Informar estado de preparación para todos los mercenarios
	r, err := c.Preparacion(ctx, &pb.PreparacionRequest{Id: int32(mercenarioID)})
	if err != nil {
		log.Fatalf("could not inform: %v", err)
	}
	if (r.Success) {
		log.Printf("Mercenario %d Preparacion: Lista", mercenarioID)
	}
	

	// Decisiones para cada piso
	if mercenarioID == 8 {
		// Interacción con el usuario
		for piso := 1; piso <= 3; piso++ {
			var decision int
			switch piso {
			case 1:
				fmt.Println("Mercenario",mercenarioID,": Ingrese su decisión para el Piso 1 (1: Escopeta, 2: Rifle automático, 3: Puños eléctricos):")
				fmt.Scan(&decision)
			case 2:
				fmt.Println("Mercenario",mercenarioID,": Ingrese su decisión para el Piso 2 (1: Pasillo A, 2: Pasillo B):")
				fmt.Scan(&decision)
			case 3:
				fmt.Println("Mercenario",mercenarioID,": Ingrese un número entre 1 y 15 para la ronda", piso, ":")
				fmt.Scan(&decision)
			}
			decisionRes, err := c.EnviarDecision(ctx, &pb.DecisionRequest{Id: int32(mercenarioID), Piso: int32(piso), Decision: int32(decision)})
			if err != nil {
				log.Fatalf("could not decide: %v", err)
			}
			log.Printf("Mercenario 8 Resultado de la decisión en Piso %d: %s", piso, decisionRes.DataNodeIp)
			if decisionRes.DataNodeIp == "Eliminado" {
				log.Printf("Mercenario 8 ha sido eliminado en el Piso %d", piso)
				os.Exit(0)
			}
		}
	} else {
		// Automatización para los otros mercenarios
		for piso := 1; piso <= 3; piso++ {
			var decision int
			switch piso {
			case 1:
				decision = rng.Intn(3) + 1 // Escoge un arma al azar
			case 2:
				decision = rng.Intn(2) + 1 // Escoge un pasillo al azar
			case 3:
				decision = rng.Intn(15) + 1 // Escoge un número entre 1 y 15
			}
			fmt.Println("Decision: ", decision, "Piso: ", piso)
			/*decisionRes, err := c.EnviarDecision(ctx, &pb.DecisionRequest{Id: int32(mercenarioID), Piso: int32(piso), Decision: int32(decision)})
			if err != nil {
				log.Fatalf("could not decide: %v", err)
			}
			log.Printf("Mercenario %d Resultado de la decisión en Piso %d: %s", mercenarioID, piso, decisionRes.DataNodeIp)
			if decisionRes.DataNodeIp == "Eliminado" {
				log.Printf("Mercenario %d ha sido eliminado en el Piso %d", mercenarioID, piso)
				return
			}*/
		}
	}
	// Mantener el programa corriendo
	select {}
}