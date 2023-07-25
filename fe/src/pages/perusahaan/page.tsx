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
  InputGroup,
  Input,
  InputRightAddon,
} from "@chakra-ui/react";
import { EditIcon, DeleteIcon, AddIcon, Search2Icon } from "@chakra-ui/icons";
import { EntityModal, useWarning } from "@/components";
import type { Perusahaan } from "@/types/models";
import { useEffect, useState } from "react";
import { useAPI } from "@/contexts";
import { useDebounce } from "usehooks-ts";
import { useSearchParams } from "react-router-dom";

const dataPerusahaan: Perusahaan[] = [
  {
    id: "1",
    nama: "Perusahaan 1",
    kode: "PRS",
    no_telp: "0231313",
    alamat: "Jl. Jalan",
  },
];

const PerusahaanPage = () => {
  const { isOpen, onClose, onOpen } = useDisclosure();
  const [isCreate, setIsCreate] = useState(false);
  const { warning, WarningModal } = useWarning();
  const [entities, setEntities] = useState<Perusahaan[]>(dataPerusahaan);
  const [entityIndex, setEntityIndex] = useState(0);
  const { api } = useAPI();
  const [rawSearchParams, setSearchParams] = useSearchParams();
  const searchParams = useDebounce(rawSearchParams, 1000);

  useEffect(() => {
    const query = {
      q: searchParams.get("q") || undefined,
    };
    api.listPerusahaan(query).then((res) => {
      if (!res.data) {
        return;
      }
      setEntities(res.data);
    });
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [searchParams]);

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
      description: "Apakah anda yakin ingin menghapus perusahaan ini?",
      onConfirm: async () => {
        const response = await api.deletePerusahaan(entities[index].id);
        if (!response.data) {
          return;
        }
        location.reload();
      },
    });
  };

  useEffect(() => {
    // TODO: API call
  }, []);

  return (
    <>
      <WarningModal />
      <EntityModal
        isOpen={isOpen}
        onClose={onClose}
        entity={entities[entityIndex] || {}}
        entityType="perusahaan"
        isCreate={isCreate}
      />
      <Flex p={8} flexDir={"column"} w="calc(100% - 250px)" maxH="100vh" overflowY="auto" gap={8}>
        <HStack gap={16}>
          <Heading w="25%">List Perusahaan</Heading>
          <Spacer />
          <InputGroup w="75%">
            <Input
              placeholder="Search..."
              onChange={(e) => {
                setSearchParams({ q: e.target.value });
              }}
            />
            <InputRightAddon>
              <IconButton w="100%" variant="unstyled" aria-label="search" icon={<Search2Icon w="100%" />} />
            </InputRightAddon>
          </InputGroup>
          <IconButton aria-label="create" icon={<AddIcon />} onClick={openCreateModal} />
        </HStack>
        <Table>
          <Thead>
            <Tr>
              <Th>Nama Perusahaan</Th>
              <Th>Kode</Th>
              <Th>No. Telp</Th>
              <Th>Alamat</Th>
              <Th>Aksi</Th>
            </Tr>
          </Thead>
          <Tbody>
            {entities.map((perusahaan, i) => {
              return (
                <Tr key={perusahaan.id}>
                  <Td>{perusahaan.nama}</Td>
                  <Td>{perusahaan.kode}</Td>
                  <Td>{perusahaan.no_telp}</Td>
                  <Td>{perusahaan.alamat}</Td>
                  <Td>
                    <HStack>
                      <IconButton aria-label="edit" icon={<EditIcon />} onClick={() => openUpdateModal(i)} />
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

export default PerusahaanPage;
