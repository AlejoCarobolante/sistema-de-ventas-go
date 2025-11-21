# Sistema de Reservas de Estacionamiento

## 📌 Descripción del Negocio
El **Sistema de Reservas de Estacionamiento** es una plataforma digital que facilita la reserva anticipada y el uso eficiente de plazas de estacionamiento (*Spot*) en diversos establecimientos (*Parking*).

### 🔹 Valor Diferencial
El negocio se sustenta en tres pilares clave:
1. **Validación Robusta**: Garantiza que un vehículo solo pueda reservar una plaza físicamente compatible, utilizando reglas explícitas de compatibilidad.
2. **Tarificación Dinámica**: Aplica costos flexibles por hora basados en el segmento de tiempo (*TimeSlot*) de la reserva.
3. **Gestión de Incumplimientos**: Maneja automáticamente las multas (*Penalty*) generadas cuando un cliente excede el tiempo de su reserva.

---

## 🧱 Modelo de Dominio

### 📋 Clases y Propósitos

| **Clase**            | **Tipo**               | **Propósito**                                                                                     |
|----------------------|------------------------|--------------------------------------------------------------------------------------------------|
| **Client**           | Entidad Principal      | Representa al usuario registrado que interactúa con la plataforma. Es el origen de la *Reservation*. |
| **Parking**          | Entidad Principal      | Representa el establecimiento físico. Contiene las *Spots*.                                      |
| **Spot**             | Entidad Principal      | Espacio individual de estacionamiento con su estado de disponibilidad.                            |
| **Reservation**      | Transacción           | Transacción central que define el compromiso de un *Client* para usar un *Spot* en un período.  |
| **Vehicle**          | Entidad Auxiliar       | Automóvil del cliente que será utilizado en la *Reservation*.                                    |
| **Payment**          | Transacción Auxiliar   | Registra la ejecución de un pago para cubrir la *Reservation* o una *Penalty*.                  |
| **Penalty**          | Transacción Auxiliar   | Modela la multa por incumplimiento (ej. exceso de tiempo) y contiene la lógica para calcular el monto. |
| **Rate**             | Configuración          | Define los costos base por tiempo y la tarifa específica para las multas (*overstayRatePerMinute*). |
| **TimeSlot**         | Configuración          | Segmento de tiempo (ej. "Hora Pico", "Fin de Semana") para aplicar tarifas variables.             |
| **VehicleType**      | Configuración          | Clasifica a los vehículos según sus características (ej. SUV, Moto) para validación.           |
| **SpotType**         | Configuración          | Clasifica a las plazas según su tamaño o funcionalidad (ej. Compacta, Estándar).                 |
| **CompatibilityRule**| Regla de Negocio       | Establece si una combinación de *VehicleType* y *SpotType* es válida para permitir la reserva.  |

---

## 🔑 Puntos Clave del Diseño

### 🔄 **Foco en la Unidireccionalidad**
- **Bajo Acoplamiento**: Las relaciones son casi siempre unidireccionales (ej. una *Reservation* conoce su *Spot*, pero el *Spot* no conoce todas sus *Reservations*).
- **Consulta por Repositorio**: Para listar todas las *Reservations* de un *Spot*, se consulta a un **Servicio de Repositorio**, no navegando directamente desde el objeto *Spot*.

### 🎯 **Separación Clara de Responsabilidades**
El modelo evita sobrecargar entidades transaccionales con lógica de negocio o configuración:

| **Concepto**               | **Clase Responsable**   | **Razón**                                                                                     |
|----------------------------|-------------------------|----------------------------------------------------------------------------------------------|
| Validación de Espacio      | *CompatibilityRule*     | La regla para el *fitting* físico se separa de *Spot* y *Vehicle*.                          |
| Cálculo de Multa           | *Penalty*               | La lógica de la multa (cálculo y registro) se aísla de *Reservation* y *Payment*.            |
| Tarificación Temporal      | *TimeSlot* y *Rate*     | La tarifa se desacopla del *Parking* y la *Reservation*, permitiendo precios variables sin modificar las entidades principales. |

---

## 📂 Estructura del Proyecto
*(Ejemplo sugerido para implementación)*

.
├── api/
│   ├── controller/          # Controladores para manejar las solicitudes HTTP
│   └── route/               # Definición de rutas
├── bootstrap/               # Inicialización de la aplicación (base de datos, configuraciones, etc.)
│   ├── app.go              # Configuración de la aplicación Gin
│   ├── database.go          # Configuración de la base de datos con GORM
│   ├── env.go               # Carga de variables de entorno
│   └── migrate.go           # Migraciones de la base de datos
├── cmd/
│   └── main.go              # Punto de entrada de la aplicación
├── domain/
├── pkg/
│   ├── constants/           # Constantes globales
│   └── usecase/             # Casos de uso (lógica de negocio)
├── .env.example             # Ejemplo de variables de entorno
├── Dockerfile               # Configuración para Docker
├── go.mod                   # Módulo de Go
├── go.sum                   # Sumas de verificación de dependencias
└── README.md                # Documentación del proyecto

# 🚗 Backend: Sistema de Reservas de Estacionamiento

**Lenguaje**: Go (Golang)
**Framework**: [Gin](https://github.com/gin-gonic/gin)
**ORM**: [GORM](https://gorm.io/)
**Base de Datos**: MySQL - MariaDB

---

## 🔹 **Características principales**
- **Validación robusta**: Reglas de compatibilidad entre vehículos y plazas.
- **Tarificación dinámica**: Precios variables según segmentos de tiempo (*TimeSlots*).
- **Gestión de multas**: Cálculo automático de penalizaciones por exceso de tiempo.
- **API RESTful**: Endpoints claros para integración con frontend o móviles.

---

## 🛠 **Tecnologías y ventajas**
| **Tecnología** | **Ventaja**                                      |
|----------------|--------------------------------------------------|
| **Go**         | Alto rendimiento, concurrencia nativa.           |
| **Gin**        | Framework minimalista y rápido para APIs.        |
| **GORM**       | ORM potente para bases de datos relacionales.    |
| **Docker**     | Despliegue consistente y escalable.              |
| **Github Actions**     | CI-CD              |