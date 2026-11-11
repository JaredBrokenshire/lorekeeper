import HelperService from "@/services/helper.service";
import api from "@/services/api";

const BASE_URL = "/campaigns"

const CampaignService = {
  list(filters) {
    const queryString = HelperService.buildSearchQueryString(filters)
    return api.get(`${BASE_URL}${queryString}`)
  }
}

export default CampaignService