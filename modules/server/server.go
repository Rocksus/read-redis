package server

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/Rocksus/read-redis/modules/messaging"
)

// Serve function returns the HTML and Javascript values to handler
func (h *Handler) Serve(w http.ResponseWriter, r *http.Request) {
	// Get user data
	users := h.UDT.GetUsers()
	// CONVERT USER DATA & REDISCOUNT INTO HTML FORMAT
	midString := `<table class="table" id="dataTable">
		<thead class="thead-dark">
			<tr>
				<th scope="col">ID</th>
				<th scope="col">Name</th>
				<th scope="col">MSISDN</th>
				<th scope="col">E-Mail</th>
				<th scope="col">Birth Date</th>
				<th scope="col">Age</th>
			</tr>
		</thead>
		<tbody>`
	frontString := `<!DOCTYPE html>
								<html lang="en">
								<head>
									<meta charset="utf-8">
									<meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
									<link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/css/bootstrap.min.css" integrity="sha384-ggOyR0iXCbMQv3Xipma34MD+dH/1fQ784/j6cY/iJTQUOhcWr7x9JvoRxT2MZw1T" crossorigin="anonymous">
								</head>
								<html>
								<nav class="navbar navbar-dark bg-primary">
								<span class="navbar-brand mb-0 h1">Visitor Count: <span id="visitorCount">` + strconv.Itoa(h.RDC.GetLatestCount()+1) + `</span></span>
						</span>
						</nav>
			<input class="form-control" id="filterData" type="text" onkeyup="filterfunc()" placeholder="Filter by name.." aria-label="Search">`
	for _, user := range users {
		midString = midString + fmt.Sprintf(`<tr>
					<th scope="row">%s</th>
					<td>%s</td>
					<td>%s</td>
					<td>%s</td>
					<td>%s</td>
					<td>%d</td>
				</tr>`, user.UserID, user.Name, user.MSISDN, user.Email, user.BirthDate.Format("01-02-2006"), user.UserAge)
	}
	midString = midString + `</tbody></table>`
	endString := `<script src="https://code.jquery.com/jquery-3.3.1.slim.min.js" integrity="sha384-q8i/X+965DzO0rT7abK41JStQIAqVgRVzpbzo5smXKp4YfRvH+8abtTE1Pi6jizo" crossorigin="anonymous"></script>
	<script src="https://cdnjs.cloudflare.com/ajax/libs/popper.js/1.14.7/umd/popper.min.js" integrity="sha384-UO2eT0CpHqdSJQ6hJty5KVphtPhzWj9WO1clHTMGa3JDZwrnQq4sF86dIHNDz0W1" crossorigin="anonymous"></script>
	<script src="https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/js/bootstrap.min.js" integrity="sha384-JjSmVgyd0p3pXB1rRibZUAYoIIy6OrQ6VrjIEaFf/nJGzIxFDsf4x0xIM+B07jRM" crossorigin="anonymous"></script>
	<script>
		function filterfunc() {
			// Declare variables
			var input, filter, table, tr, td, i, txtValue;
			input = document.getElementById("filterData");
			filter = input.value.toUpperCase();
			table = document.getElementById("dataTable");
			tr = table.getElementsByTagName("tr");

			// Loop through all table rows, and hide those who don't match the search query
			for (i = 1; i < tr.length; i++) {
				td = tr[i].getElementsByTagName("td")[0];
				if (td) {
				txtValue = td.textContent || td.innerText;
				if (txtValue.toUpperCase().indexOf(filter) > -1) {
					tr[i].style.display = "";
				} else {
					tr[i].style.display = "none";
				}
				}
			}
		}
	</script>
	</body></html>`
	// WRITE w HTTP SERVER WITH HTML
	fmt.Fprintf(w, fmt.Sprintf("%s%s%s", frontString, midString, endString))

	// SEND VISITOR
	// initiate producer
	prodConf := messaging.ProducerConfig{
		NsqdAddress: os.Getenv("nsqdAddr"),
	}
	prod := messaging.NewProducer(prodConf)

	// publish message
	topic := os.Getenv("nsqTopic")
	msg := "1"
	prod.Publish(topic, msg)
}
