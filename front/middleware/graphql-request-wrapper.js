export default function(axios, url, data) {
  return axios.post(url, data).then(function(response) {
    if ('errors' in response.data) return Promise.reject(response.data.errors)
    else return response.data.data
  })
}
