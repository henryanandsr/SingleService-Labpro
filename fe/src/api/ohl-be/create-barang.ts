import { Barang, Response } from "@/types/models"
import {api} from "."

type CreateBarangData = Barang

const createBarangAPI = async (barang: Barang) => {
  // eslint-disable-next-line @typescript-eslint/no-unused-vars
  const {id, ...rest}  = barang;
  const response = await api.post<Response<CreateBarangData>>("/barang", {
      ...rest
  })
  return response.data
}

export default createBarangAPI;