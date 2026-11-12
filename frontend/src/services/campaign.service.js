import api from "@/services/api";
import HelperService from "@/services/helper.service";

const BASE_URL = "/campaigns"

const CampaignService = {
  list(filters) {
    const queryString = HelperService.buildSearchQueryString(filters)
    return api.get(`${BASE_URL}${queryString}`)
  },
  get(id) {
    return api.get(`${BASE_URL}/${id}`)
  },
  uploadSessionNotes(id, file) {
    const formData = new FormData();
    formData.append("file", file)

    return api.post(`${BASE_URL}/${id}/files`, formData, {
      headers: {
        "Content-Type": "multipart/form-data",
      }
    })
  }
}

export default CampaignService