import {  Perusahaan, Response } from "@/types/models";
import {api} from "."


interface listPerusahaanQuery {
    q?: string;
}

type ListPerusahaanData = Perusahaan[]

const listPerusahaanAPI = async (query: listPerusahaanQuery) => {
    const response = await api.get<Response<ListPerusahaanData>>("/perusahaan", {
        params: query
    })
    return response.data
} 

export default listPerusahaanAPI;