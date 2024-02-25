import axios from "axios";
const API_URL = "http://localhost:8080";
const API_CODESANCUS = async (urlObject, payload, type) => {
  return await axios({
    method: urlObject.method,
    url:`${API_URL}/${urlObject.endpoint}/${type}`,
    data:payload,
    headers:{"Content-Type":"application/json"}
  });
};
export default API_CODESANCUS;
