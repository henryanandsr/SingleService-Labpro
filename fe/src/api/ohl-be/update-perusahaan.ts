import { Perusahaan, Response } from "@/types/models"
import {api} from "."

type UpdatePerusahaanData = Perusahaan

const updatePerusahaanAPI = async (perusahaan: Perusahaan) => {
    const {id: perusahaan_id, ...rest} = perusahaan
    const response = await api.put<Response<UpdatePerusahaanData>>(`/perusahaan/${perusahaan_id}`, {
        ...rest
    })
    return response.data
}

export default updatePerusahaanAPI;