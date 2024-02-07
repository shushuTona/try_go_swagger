import { Api } from './gen/api'

const apiConfig = {
    baseUrl: "http://localhost:8081/api"
}
const api = new Api( apiConfig )

export { api }
