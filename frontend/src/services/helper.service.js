import axios from "@axios"

class HelperService {
    // Axios Instance
    axiosIns = null

    constructor(axiosIns) {
      this.axiosIns = axiosIns
    }

    /**
     * Build a query string from a filters object using key value pairs
     * */
    buildSearchQueryString(filters) {
      let query = '?';
      for (const key in filters) {
        if (!filters[key]) {
          continue
        }
        if (query.length > 1) {
          query += "&"
        }
        query += `${key}=${filters[key]}`;
      }
      return query;
    }

    /**
     * Get status code and message from error response
     * */
    getErrorInfo(error) {
      const statusCode = error?.response?.status;
      const msg = error?.response?.data?.error;
      return {
        statusCode,
        msg,
      }
    }

    /**
     * Create an error toast notification
     * */
    errorToast(toast, error, customMessage) {
      const {msg} = this.getErrorInfo(error)
      if (msg) {
        toast.error(msg, {
          toastClassName: ['toast-std', 'warning-toast'],
        })
        return msg;
      }

      toast.error(customMessage, {
        toastClassName: ['toast-std', 'warning-toast'],
      })
      return customMessage;
    }

    /**
     * Create a success toast notification
     * */
    successToast(toast, customMessage) {
      toast.success(customMessage, {
        toastClassName: ['toast-std', 'success-toast'],
      })
      return customMessage;
    }

    /**
     * Return the API URL
     * */
    getApiUrl() {
      return this.axiosIns.defaults.baseURL
    }
}

const helperService = new HelperService(axios)
export default helperService