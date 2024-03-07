package zoom

type Status string

const (
	StatusArrivedInDestination Status = "ARRIVED_IN_DESTINATION"
	StatusAvailableForPickup   Status = "AVAILABLE_FOR_PICKUP"
	StatusDelivered            Status = "DELIVERED"
)

var statusMap = map[string]Status{
	"RECIBIDO EN LA TAQUILLA DESTINO":       StatusArrivedInDestination,
	"DISPONIBLE PARA EL RETIRO EN TAQUILLA": StatusAvailableForPickup,
	"ENTREGADO AL CLIENTE":                  StatusDelivered,
	"ENTREGADO AL DESTINO":                  StatusDelivered,
}

type GetTrackingResponse struct {
	Codrespuesta string   `json:"codrespuesta"`
	Mensaje      string   `json:"mensaje"`
	Shipment     Shipment `json:"entidadRespuesta"`
}

type Shipment struct {
	InfoDhl  string     `json:"infoDhl"`
	InfoZoom InfoZoom   `json:"infoZoom"`
	Tracking []Tracking `json:"tracking"`
	Tipocon  string     `json:"tipocon"`
	InfoGuia string     `json:"infoGuia"`
}

func (s *Shipment) GetStatus() (isFinal bool, status Status) {
	status, isFinal = statusMap[s.InfoZoom.DescripcionEstatus]
	if !isFinal {
		status, isFinal = statusMap[s.Tracking[0].Estatus.Nombre]
		if !isFinal {
			return
		}
	}
	return false, Status(s.InfoZoom.DescripcionEstatus)
}

type InfoZoom struct {
	Codguia            int64          `json:"codguia"`
	Referencia         string         `json:"referencia"`
	Tipoenvio          string         `json:"tipoenvio"`
	Fecha              string         `json:"fecha"`
	Origen             Origen         `json:"origen"`
	Destino            Destino        `json:"destino"`
	Oficinaorigen      Oficinaorigen  `json:"oficinaorigen"`
	Oficinadestino     Oficinadestino `json:"oficinadestino"`
	Codoficinaori      int64          `json:"codoficinaori"`
	OficinaDestino     string         `json:"oficina_destino"`
	Codestatus         int64          `json:"codestatus"`
	DescripcionEstatus string         `json:"descripcion_estatus"`
	Mensajeweb         string         `json:"mensajeweb"`
	Codservicio        int64          `json:"codservicio"`
	Nombreservicio     string         `json:"nombreservicio"`
	Icono              interface{}    `json:"icono"`
	Color              interface{}    `json:"color"`
	Peso               string         `json:"peso"`
	Nropiezas          int64          `json:"nropiezas"`
	Coddestinatario    interface{}    `json:"coddestinatario"`
	Guiadhl            string         `json:"guiadhl"`
	Codcasillero       string         `json:"codcasillero"`
}

type Destino struct {
	Codciudad int64  `json:"codciudad"`
	Nombre    string `json:"nombre"`
	Codestado int64  `json:"codestado"`
	Estado    Estado `json:"estado"`
}

type Estado struct {
	Codestado int64  `json:"codestado"`
	Nombre    string `json:"nombre"`
	Codpais   int64  `json:"codpais"`
}

type Oficinadestino struct {
	Codoficina int64  `json:"codoficina"`
	Nombre     string `json:"nombre"`
	Codciudad  int64  `json:"codciudad"`
	Siglas     string `json:"siglas"`
	Codcliente string `json:"codcliente"`
	Codtipoofi int64  `json:"codtipoofi"`
	Direccion  string `json:"direccion"`
	Ciudad     Origen `json:"ciudad"`
}

type Origen struct {
	Codciudad        int64   `json:"codciudad"`
	Nombre           string  `json:"nombre"`
	Codoficinaope    int64   `json:"codoficinaope"`
	Oficinaoperativa Oficina `json:"oficinaoperativa"`
}

type Oficina struct {
	Codoficina int64   `json:"codoficina"`
	Nombre     string  `json:"nombre"`
	Siglas     string  `json:"siglas"`
	Direccion  *string `json:"direccion,omitempty"`
}

type Oficinaorigen struct {
	Codoficina int64  `json:"codoficina"`
	Nombre     string `json:"nombre"`
	Codciudad  int64  `json:"codciudad"`
	Siglas     string `json:"siglas"`
	Ciudad     Origen `json:"ciudad"`
}

type Tracking struct {
	ID            int64       `json:"id"`
	Receptor      string      `json:"receptor"`
	Codtracking   interface{} `json:"codtracking"`
	Codestatus    int64       `json:"codestatus"`
	Observacion   string      `json:"observacion"`
	Codusuario    int64       `json:"codusuario"`
	Horareal      string      `json:"horareal"`
	Codmensajero  *int64      `json:"codmensajero"`
	Codruta       int64       `json:"codruta"`
	Codoficinaori int64       `json:"codoficinaori"`
	Fechapro      string      `json:"fechapro"`
	Sello         string      `json:"sello"`
	Hora          *string     `json:"hora"`
	Fechahorareal string      `json:"fechahorareal"`
	Estatus       Estatus     `json:"estatus"`
	Usuario       Usuario     `json:"usuario"`
	Ruta          *Ruta       `json:"ruta"`
	Oficina       Oficina     `json:"oficina"`
	Mensajero     *Mensajero  `json:"mensajero"`
	Track         string      `json:"track"`
}

type Estatus struct {
	Codestatus int64  `json:"codestatus"`
	Nombre     string `json:"nombre"`
	Siglas     string `json:"siglas"`
	Nombreweb  string `json:"nombreweb"`
}

type Mensajero struct {
	Codmensajero int64  `json:"codmensajero"`
	Nombre       string `json:"nombre"`
}

type Ruta struct {
	Codruta       int64      `json:"codruta"`
	Nombre        string     `json:"nombre"`
	Codoficinaret int64      `json:"codoficinaret"`
	Numeroruta    int64      `json:"numeroruta"`
	Oficinaret    Oficinaret `json:"oficinaret"`
}

type Oficinaret struct {
	Codoficina int64  `json:"codoficina"`
	Nombre     string `json:"nombre"`
	Codcliente string `json:"codcliente"`
}

type Usuario struct {
	Codusuario int64  `json:"codusuario"`
	Nombre     string `json:"nombre"`
}
