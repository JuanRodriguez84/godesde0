package modelos_estructuras

type Hombre struct {
	edad int
	altura float32
	peso float32
	respirando bool
	pensando bool
	comiendo bool
	vivo bool
}

// para que el hombre implemente la interface de humano, para reconocer al Hombre como un humano, el Hombre debe tener las 
// mismas funciones que el humano
func (h *Hombre) Respirar(){
	h.respirando = true
}

func (h *Hombre) Comer(){
	h.comiendo = true
}

func (h *Hombre) Pensar(){
	h.pensando = true
}
func (h *Hombre) Sexo() string{
	return "Hombre"
}

func (h *Hombre) EstaVivo() bool{
	return true
}

// Go con solo encontrar las declaraciones (Respirar(), Comer(), Pensar(), Sexo()) de una interface asume que debe implementar sin necesidad de
// colocar implements 
