## Contexto del Proyecto

### Problema y necesidad de la empresa
Una empresa de desarrollo de soluciones tecnológicas ha inaugurado una segunda sede operativa y enfrenta el reto de mantener una comunicación estructurada, segura y privada entre los empleados de ambas sedes. Actualmente, necesitan una plataforma que opere internamente sin depender de servidores públicos (SMTP externo), garantizando el control total de la información institucional.

### Selección del sistema y Justificación
* **Sistema seleccionado:** Gestión de correos internos.
* **Justificación (Arquitectura Cliente-Servidor):** Se seleccionó este sistema porque resuelve directamente la necesidad de comunicación formal intraempresarial. Su arquitectura permite implementar operaciones cliente-servidor reales: un cliente (ubicado en una sede) consume la funcionalidad de enviar, leer o eliminar correos, y el servidor (ubicado en otra VLAN o sede) procesa la solicitud, almacena la información y emite una respuesta tras validar la sesión del usuario.

### Datos medibles para Estadística
Este sistema generará métricas cuantitativas que permitirán aplicar estadística descriptiva e inferencial. Se podrán medir variables como:
* Tiempos de respuesta del servidor (latencia).
* Cantidad de mensajes enviados por sede.
* Frecuencia de uso por usuario.
* Tasa de errores de autenticación.

### Actores y usuarios iniciales
* Empleados de Sede A (usuarios finales).
* Empleados de Sede B (usuarios finales).
* Administradores de Red/Sistema.

---

## Cronograma de Trabajo

El proyecto se desarrollará en 16 semanas divididas en tres parciales:

| Fase | Periodo | Actividades Principales |
| :--- | :--- | :--- |
| **Parcial 1** | Semanas 1 - 5 | Definición del contexto, planificación, especificación de requisitos y análisis socio-tecnológico. |
| **Parcial 2** | Semanas 6 - 10 | Diseño de la arquitectura (red y software) y desarrollo del prototipo inicial. |
| **Parcial 3** | Semanas 11 - 16 | Integración, despliegue, pruebas estadísticas, presentación y cierre definitivo. |

---

## Declaración de Integridad Académica, Licencia y Uso de IA
> *"El equipo certifica que el código y configuraciones presentadas son de autoría propia. El proyecto se distribuirá bajo una licencia Open Source. Declaramos que cualquier uso de herramientas de Inteligencia Artificial será utilizado exclusivamente como apoyo al desarrollo, citando debidamente su intervención, respetando los principios de integridad académica e intelectual de la institución".*