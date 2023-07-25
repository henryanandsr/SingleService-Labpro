import {
  Flex,
  Heading,
  Table,
  Thead,
  Tr,
  Th,
  Td,
  Tbody,
  IconButton,
  HStack,
  useDisclosure,
  Spacer,
  Input,
  InputGroup,
  InputRightAddon,
  Select,
} from "@chakra-ui/react";
import { EditIcon, DeleteIcon, AddIcon, Search2Icon } from "@chakra-ui/icons";
import { EntityModal, useWarning } from "@/components";
import type {
  BarangWithPerusahaan as Barang,
  Perusahaan,
} from "@/types/models";
import { useEffect, useState } from "react";
import { useAPI } from "@/contexts";
import { useSearchParams } from "react-router-dom";
import { useDebounce } from "usehooks-ts";

const BarangPage = () => {
  const { isOpen, onClose, onOpen } = useDisclosure();
  const [isCreate, setIsCreate] = useState(false);
  const { warning, WarningModal } = useWarning();
  const [listPerusahaan, setListPerusahaan] = useState<Perusahaan[]>([]);
  const [entities, setEntities] = useState<Barang[]>([]);
  const [entityIndex, setEntityIndex] = useState(0);
  const [rawSearchParams, setSearchParams] = useSearchParams({});
  const searchParams = useDebounce(rawSearchParams, 1000);
  const { api } = useAPI();

  const fetch = async () => {
    const query = {
      q: searchParams.get("q") || undefined,
      perusahaan: searchParams.get("perusahaan") || undefined,
    };
    const res = await api.listBarang(query);
    if (!res.data) {
      return;
    }

    const barangWithPerush = await Promise.all(res.data.map(async (e) => {
      const perushResp = await api.detailPerusahaan(e.perusahaan_id);
      const perusahaan = perushResp.data ? perushResp.data : { nama: ""}

      return {
        ...e,
        perusahaan,
      };
    }))

    setEntities(barangWithPerush);
  };

  useEffect(() => {
    fetch();
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [searchParams]);

  useEffect(() => {
    api.listPerusahaan({}).then((res) => {
      if (!res.data) {
        return;
      }
      setListPerusahaan(res.data);
    });
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  const openCreateModal = () => {
    setIsCreate(true);
    onOpen();
  };

  const openUpdateModal = (index: number) => {
    setEntityIndex(index);
    setIsCreate(false);
    onOpen();
  };

  const handleDelete = (index: number) => {
    warning({
      title: "Hapus Barang",
      description: "Apakah anda yakin ingin menghapus barang ini?",
      onConfirm: async () => {
        const response = await api.deleteBarang(entities[index].id);
        if (!response.data) {
          return;
        }
        // refresh page
        location.reload();
      },
    });
  };

  return (
    <>
      <WarningModal />
      <EntityModal
        isOpen={isOpen}
        onClose={onClose}
        entity={entities[entityIndex] || {}}
        entityType="barang"
        isCreate={isCreate}
        selectChoices={{
          perusahaan_id: {
            ...listPerusahaan.reduce((acc, cur) => {
              acc[cur.id] = cur.nama;
              return acc;
            }, {} as Record<string, string>),
          },
        }}
      />
      <Flex
        p={8}
        flexDir={"column"}
        w="calc(100% - 250px)"
        maxH="100vh"
        overflowY="auto"
        gap={8}
      >
        <HStack gap={16}>
          <Heading w="25%">List Barang</Heading>
          <Select
            w="25%"
            onChange={(e) => {
              setSearchParams({ perusahaan: e.target.value });
            }}
          >
            <option value="">Semua Perusahaan</option>
            {listPerusahaan.map((p) => (
              <option key={p.id} value={p.id}>
                {p.nama}
              </option>
            ))}
          </Select>
          <InputGroup w="50%">
            <Input
              placeholder="Search..."
              onChange={(e) => {
                setSearchParams({ q: e.target.value });
              }}
            />
            <InputRightAddon>
              <IconButton
                w="100%"
                variant="unstyled"
                aria-label="search"
                icon={<Search2Icon w="100%" />}
              />
            </InputRightAddon>
          </InputGroup>
          <Spacer />
          <IconButton
            aria-label="create"
            icon={<AddIcon />}
            onClick={openCreateModal}
          />
        </HStack>
        <Table>
          <Thead>
            <Tr>
              <Th>Nama Barang</Th>
              <Th>Kode</Th>
              <Th>Stok</Th>
              <Th>Harga</Th>
              <Th>Perusahaan</Th>
              <Th>Aksi</Th>
            </Tr>
          </Thead>
          <Tbody>
            {entities.map((barang, i) => {
              return (
                <Tr key={barang.id}>
                  <Td>{barang.nama}</Td>
                  <Td>{barang.kode}</Td>
                  <Td>{barang.stok}</Td>
                  <Td>{barang.harga}</Td>
                  <Td>{barang.perusahaan.nama}</Td>
                  <Td>
                    <HStack>
                      <IconButton
                        aria-label="edit"
                        icon={<EditIcon />}
                        onClick={() => openUpdateModal(i)}
                      />
                      <IconButton
                        aria-label="delete"
                        icon={<DeleteIcon />}
                        colorScheme="red"
                        onClick={() => handleDelete(i)}
                      />
                    </HStack>
                  </Td>
                </Tr>
              );
            })}
          </Tbody>
        </Table>
      </Flex>
    </>
  );
};

export default BarangPage;
