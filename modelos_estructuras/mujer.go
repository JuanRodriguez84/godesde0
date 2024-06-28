package modelos_estructuras

type Mujer struct {
	Hombre           // hereda del hombre como en POO
	EsMadre bool     // hereda lo del hombre y tiene propiedades nuevas
}

// para que el hombre implemente la interface de humano, para reconocer al Hombre como un humano, el Hombre debe tener las 
// mismas funciones que el humano
func (m *Mujer) Respirar(){
	m.respirando = true
}

func (m *Mujer) Comer(){
	m.comiendo = true
}

func (m *Mujer) Pensar(){
	m.pensando = true
}
func (m *Mujer) Sexo() string{
	return "Mujer"
}

func (m *Mujer) EstaVivo() bool{
	return true
}

// Go con solo encontrar las declaraciones (Respirar(), Comer(), Pensar(), Sexo()) de una interface asume que debe implementar sin necesidad de
// colocar implements 
