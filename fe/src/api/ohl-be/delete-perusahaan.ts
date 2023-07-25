import { Perusahaan, Response } from "@/types/models"
import {api} from "."

type DeletePerusahaanData = Perusahaan

const deletePerusahaanAPI= async (perusahaan_id: string) => {
    const response = await api.delete<Response<DeletePerusahaanData>>(`/perusahaan/${perusahaan_id}`)
    return response.data
}

export default deletePerusahaanAPI;