import { Perusahaan, Response } from "@/types/models"
import {api} from "."

type DetailPerusahaanData = Perusahaan

const detailPerusahaanAPI = async (perusahaan_id: string) => {
    const response = await api.get<Response<DetailPerusahaanData>>(`/perusahaan/${perusahaan_id}`)
    return response.data
}

export default detailPerusahaanAPI;