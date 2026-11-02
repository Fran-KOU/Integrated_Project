| Integrante | Rol en el Proyecto | Responsabilidades Principales |
| :--- | :--- | :--- |
| **David Torres** | Desarrollador Backend (Go) | Desarrollo de la aplicación empresarial cliente-servidor en Go, aplicando modularidad y manejo de errores. Diseño de las estructuras (structs) y la lógica de la bandeja de entrada, lectura y envío de mensajes. |
| **Francis Acuña** | Administrador de Infraestructura y Redes | Diseño e implementación de la topología distribuida con OpenWrt. Configuración de direccionamiento VLSM, VLANs, OSPF, RSTP, enlaces troncales y segmentación segura. |
| **Iván Gualotuña** | Especialista en Seguridad y Analista de Datos | Desarrollo del servicio de autenticación en FastAPI (registro, inicio de sesión y protección de contraseñas). Diseño del instrumento de recolección de datos, análisis estadístico y redacción del impacto socio-tecnológico. |

| Actor | Tipo | Qué puede hacer |
| :--- | :--- | :--- |
| **Empleado** | Persona | Iniciar sesión, enviar mensajes, consultar su bandeja de entrada y leer sus mensajes. No puede ver buzones ajenos ni administrar cuentas. |
| **Administrador** | Persona | Iniciar sesión y desbloquear cuentas bloqueadas por intentos fallidos. No puede leer mensajes de otros usuarios. |
| **Servicio de autenticación (FastAPI)** | Servicio externo | Valida usuario y contraseña, emite el token de sesión, cuenta los intentos fallidos y bloquea o desbloquea cuentas. |
| **Servidor de correos (Go)** | Servicio externo | Guarda los buzones, recibe los mensajes enviados y entrega la bandeja, siempre que la petición lleve un token válido. |