import { Barang, Response } from "@/types/models"
import {api} from "."

type DeleteBarangData = Barang

const deleteBarangAPI = async (barang_id: string) => {
    const response = await api.delete<Response<DeleteBarangData>>(`/barang/${barang_id}`)
    return response.data
}

export default deleteBarangAPI;