package lab3

import (
	"log"
	"golang.org/x/net/context"
)

type Server struct {
	hay_info bool
	planeta string
	ciudad string
	soldados int32
}

func (s *Server) Enviarinfo(ctx context.Context, in *Info) (*Info, error){
	if in.Planeta != "" {
		log.Printf("Hay  informacion, enviar a broker: %s\n", in.Planeta)
		s.hay_info = true
		s.planeta = in.Planeta
		s.ciudad = in.Ciudad
		s.soldados = in.Soldados
	}
	return &Info{Planeta: "Tatooine",Ciudad: "Mos_Eisley", Soldados: 5},nil
}

func (s *Server) Alertabroken(ctx context.Context, in *Message) (*Info, error){
	if in.Inicio == "Inicio"{
		s.hay_info = false
	}
	for(s.hay_info != true){
	}
	return &Info{Planeta: s.planeta,Ciudad: s.ciudad,Soldados: s.soldados},nil
}