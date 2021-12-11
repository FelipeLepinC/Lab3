package lab3

import (
	"log"
	"golang.org/x/net/context"
	"math/rand"
	"time"
)

type Server struct {
	hay_info bool
	planeta string
	ciudad string
	soldados int32
}

func (s *Server) Enviarinfo(ctx context.Context, in *Info) (*Info, error){
	return &Info{Planeta: "Tatooine",Ciudad: "Mos_Eisley", Soldados: 5},nil
}

func (s *Server) Fulcrum(ctx context.Context, in *Operacion) (*Reloj, error){
	log.Printf("%s %s %s %d", in.Accion, in.Planeta,in.Ciudad,in.Valor)
	return &Reloj{Planeta: in.Planeta,X:1,Y:0,Z:0},nil
}

func (s *Server) Alertabroken(ctx context.Context, in *Operacion) (*Message, error){
	n := int32(0)
	if in.Accion != "" {
		rand.Seed(time.Now().UnixNano())
		n = int32(1 + rand.Intn(3-1+1))
		log.Printf("Servidor elegido es : %d",n)
	}
	return &Message{Nserver:n},nil
}