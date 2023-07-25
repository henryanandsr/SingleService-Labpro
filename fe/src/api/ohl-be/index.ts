import axios from "axios"
import loginAPI from "./login"
import selfAPI from "./self"
import listBarangAPI from "./list-barang"
import createBarangAPI from "./create-barang"
import updateBarangAPI from "./update-barang"
import deleteBarangAPI from "./delete-barang"
import detailBarangAPI from "./detail-barang"
import listPerusahaanAPI from "./list-perusahaan"
import createPerusahaanAPI from "./create-perusahaan"
import updatePerusahaanAPI from "./update-perusahaan"
import deletePerusahaanAPI from "./delete-perusahaan"
import detailPerusahaanAPI from "./detail-perusahaan"

export const api = axios.create()

const OHLBEAPI = {
  // Auth
  login: loginAPI,
  self: selfAPI,

  // Barang
  listBarang: listBarangAPI,
  createBarang: createBarangAPI,
  updateBarang: updateBarangAPI,
  deleteBarang: deleteBarangAPI,
  detailBarang: detailBarangAPI,

  // Perusahaan
  listPerusahaan: listPerusahaanAPI,
  createPerusahaan: createPerusahaanAPI,
  updatePerusahaan: updatePerusahaanAPI,
  deletePerusahaan: deletePerusahaanAPI,
  detailPerusahaan: detailPerusahaanAPI
}

// add try catch for every key
for (const key in OHLBEAPI) {
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  const fn = OHLBEAPI[key as keyof typeof OHLBEAPI] as any
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  OHLBEAPI[key as keyof typeof OHLBEAPI] = async (...args: any[]) => {
    try {
      const res = await fn(...args)
      if (res.status === "error") throw new Error(res.message)
      return res
    } catch (error) {
      console.error(error) // Untuk debugging 😁
      if (error instanceof Error){
        return {
          status: "error",
          message: error.message,
          data: null
        }
      } else {
        return {
          status: "error",
          message: "Unknown error",
          data: null
        }
      }
    }
  }
}

export type OHLBEAPIType = typeof OHLBEAPI
export default OHLBEAPI