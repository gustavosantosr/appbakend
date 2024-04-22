package middlew

import(
	"net/http"
	"github.com/gustavosantosr/twittor/bd"
)
/*ChequeoBD middleware que permite conocer el estado de la conexion*/
func ChequeoBD(next http.HandlerFunc) http.HandlerFunc{
	return func( w http.ResponseWriter, r *http.Request){
		if bd.ChequeoConnection()==0{
			http.Error(w, "Conexion perdida con la base de datos", 500)
			return
		}
		next.ServeHTTP(w,r)
	}
}