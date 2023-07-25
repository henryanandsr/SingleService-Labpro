import {  Perusahaan, Response } from "@/types/models"
import {api} from "."

type CreatePerusahaanData = Perusahaan

const createPerusahaanAPI = async (perusahaan: Perusahaan) => {
    const response = await api.post<Response<CreatePerusahaanData>>("/perusahaan", {
        ...perusahaan
    })
    return response.data
}

export default createPerusahaanAPI;