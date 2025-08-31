package report

type ReportModel struct{
	NombreArchivo string `json:"nombreArchivo"`
	Ruta string `json:"rutaArchivo"`
	Size uint64 `json:"sizeArchivo"`
	Hash string `json:"hashArchivo"`
}