import api from "@/services/api";

const BASE_URL = "/files"

const FileService = {
  get(fileLocation) {
    return api.get(`${BASE_URL}/${fileLocation}`, { responseType: 'blob' })
  }
}

export default FileService