export type User = {
  username: string;
}

export type Barang = {
  id: string;
  nama: string;
  kode: string;
  harga: number;
  stok: number;
  perusahaan_id: string;
}

export type BarangWithPerusahaan = Barang & {
  perusahaan: {
    nama: string;
  }
}

export type Perusahaan = {
  id: string;
  nama: string;
  kode: string;
  alamat: string;
  no_telp: string;
}

export type Response<T> = {
  status: "success" | "error";
  message: string;
  data: T | null;
}