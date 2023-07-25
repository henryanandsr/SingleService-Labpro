import { Barang, Response } from "@/types/models"
import {api} from "."

type DetailBarangData = Barang

const detailBarangAPI = async (barang_id: string) => {
    const response = await api.get<Response<DetailBarangData>>(`/barang/${barang_id}`)
    return response.data
}

export default detailBarangAPI;