import { Barang, Response } from "@/types/models"
import {api} from "."

type UpdateBarangData = Barang

const updateBarangAPI = async (barang: Barang) => {
  const {id: barang_id, ...rest} = barang
  const response = await api.put<Response<UpdateBarangData>>(`/barang/${barang_id}`, {
      ...rest
  })
  return response.data
}

export default updateBarangAPI;