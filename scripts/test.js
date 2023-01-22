import http from 'k6/http'

export let options = {
   vus: 5,
   dulation: '5s'
}

export default function() {
   http.get("http://host.docker.internal:3500/ht")
}