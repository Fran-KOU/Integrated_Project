package main

// INTERFAZ

// Transporte define los métodos de comunicación en red
type Transporte interface {
	Leer() []byte
	Escribir(datos []byte) error
	Cerrar() error
}

// ESTRUCTURAS BASE

type SesionUsuario struct {
	idUsuario       string
	direccionCorreo string
	token           string
}

func (s *SesionUsuario) EsValida() bool {
	// Lógica para validar el token con FastAPI
	return true
}

type Correo struct {
	id           string
	remitente    string
	destinatario string
	asunto       string
	cuerpo       string
	marcaTiempo  int64
}

func (c *Correo) A_JSON() []byte {
	// Lógica para serializar el correo a JSON
	return []byte{}
}

func (c *Correo) Desde_JSON(datos []byte) error {
	// Lógica para deserializar el JSON a la estructura Correo
	return nil
}

type Buzon struct {
	idPropietario  string
	bandejaEntrada []Correo // Agregación de Correo
	enviados       []Correo // Agregación de Correo
}

func (b *Buzon) AgregarCorreoRecibido(correo Correo) {
	b.bandejaEntrada = append(b.bandejaEntrada, correo)
}

func (b *Buzon) AgregarCorreoEnviado(correo Correo) {
	b.enviados = append(b.enviados, correo)
}

func (b *Buzon) ObtenerBandejaEntrada() []Correo {
	return b.bandejaEntrada
}

// LADO DEL SERVIDOR

type ManejadorCliente struct {
	conexion Transporte    // Implementa la interfaz Transporte
	sesion   SesionUsuario // Agregación de la sesión actual
}

func (m *ManejadorCliente) EscucharPeticiones() {
	// Bucle para leer datos de la conexión
}

func (m *ManejadorCliente) Desconectar() {
	// Lógica para cerrar la conexión
}

type ServidorCorreo struct {
	direccion          string
	conexionesClientes map[string]ManejadorCliente // Composición (0..*)
	buzones            map[string]Buzon            // Composición (0..*)
}

func (s *ServidorCorreo) Iniciar() error {
	// Lógica para levantar el servidor y escuchar conexiones
	return nil
}

func (s *ServidorCorreo) EnrutarCorreo(correo Correo) error {
	// Lógica para buscar el buzón destino y entregar el correo
	return nil
}

func (s *ServidorCorreo) Autenticar(token string) bool {
	// Lógica de verificación
	return true
}

// LADO DEL CLIENTE

type ClienteCorreo struct {
	ipServidor   string
	conexion     Transporte    // Implementa la interfaz Transporte
	sesionActual SesionUsuario // Composición (1)
}

func (c *ClienteCorreo) IniciarSesion(correo string, contrasena string) error {
	// Lógica para obtener el token y establecer la sesionActual
	return nil
}

func (c *ClienteCorreo) EnviarCorreo(destinatario string, asunto string, cuerpo string) error {
	// Lógica para construir el struct Correo, serializar a JSON y enviar por la conexión
	return nil
}

func (c *ClienteCorreo) ObtenerBandeja() []Correo {
	// Lógica para solicitar la bandeja al servidor y retornar los correos
	return []Correo{}
}