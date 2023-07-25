import { useAPI } from "@/contexts";
import { Barang, Perusahaan } from "@/types/models";
import {
  Button,
  Container,
  Input,
  Modal,
  ModalBody,
  ModalContent,
  ModalFooter,
  ModalHeader,
  ModalOverlay,
  Select,
  Text,
} from "@chakra-ui/react";
import { useEffect, useState } from "react";

interface EntityModalProps {
  isOpen: boolean;
  entity: Record<string, string | number | Record<string, string>>;
  entityType: "barang" | "perusahaan";
  isCreate: boolean;
  onClose: () => void;
  selectChoices?: Record<string, Record<string, string>>;
}

const EntityModal = ({ isOpen, entity, entityType, isCreate, onClose, selectChoices }: EntityModalProps) => {
  const { api } = useAPI();
  const [editableEntity, setEditableEntity] = useState<Record<string, string | number | Record<string, string>>>(
    entity || {}
  );

  const handleSave = async () => {
    if (entityType === "barang") {
      if (isCreate) {
        const res = await api.createBarang(editableEntity as Barang);
        if (!res.data) return;
        location.reload();
      } else {
        const res = await api.updateBarang(editableEntity as Barang);
        if (!res.data) return;
        location.reload();
      }
    } else {
      if (isCreate) {
        const res = await api.createPerusahaan(editableEntity as Perusahaan);
        if (!res.data) return;
        location.reload();
      } else {
        const res = await api.updatePerusahaan(editableEntity as Perusahaan);
        if (!res.data) return;
        location.reload();
      }
    }
    onClose();
  };

  useEffect(() => {
    if (isCreate) {
      // copy keys into new object
      const newEntity: Record<string, string | number> = {};
      if (Object.keys(entity).length > 0) {
        Object.keys(entity).forEach((key) => {
          if (entity[key] instanceof Object) return;
          if (key.includes("id")) return;
          newEntity[key] = typeof entity[key] === "string" ? "" : 0;
          if (entityType === "barang" && selectChoices?.perusahaan_id)
            newEntity["perusahaan_id"] = Object.keys(selectChoices.perusahaan_id)[0];
        });
      } else {
        if (entityType === "barang") {
          newEntity["nama"] = "";
          newEntity["kode"] = "";
          newEntity["stok"] = 0;
          newEntity["harga"] = 0;
          newEntity["perusahaan_id"] = selectChoices?.perusahaan_id
            ? Object.values(selectChoices?.perusahaan_id)[0]
            : "";
        } else {
          newEntity["nama"] = "";
          newEntity["kode"] = "";
          newEntity["no_telp"] = "";
          newEntity["alamat"] = "";
        }
      }
      setEditableEntity(newEntity);
    } else {
      setEditableEntity(entity);
    }
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [isCreate, isOpen]);

  return (
    <Modal isOpen={isOpen} onClose={onClose}>
      <ModalOverlay />
      <ModalContent>
        <ModalHeader>{`${isCreate ? "Create" : "Edit"} ${entityType}`}</ModalHeader>
        <ModalBody>
          {Object.entries(editableEntity).map(([key, value]) => {
            if (value instanceof Object) return;
            if (key.includes("id")) return;
            return (
              <Container key={key}>
                <Text mt={2}>{key.charAt(0).toUpperCase() + key.slice(1)}</Text>
                <Input
                  value={value}
                  required
                  type={typeof value === "number" ? "number" : "text"}
                  onChange={(e) => {
                    setEditableEntity({
                      ...editableEntity,
                      [key]: typeof value === "number" ? e.target.valueAsNumber ?? 0 : e.target.value,
                    });
                  }}
                />
              </Container>
            );
          })}
          {Object.entries(editableEntity).map(([key, _]) => {
            if (Object.keys(selectChoices || {}).includes(key)) {
              return (
                <Container key={key}>
                  <Text mt={2}>{key.charAt(0).toUpperCase() + key.split("_")[0].slice(1)}</Text>
                  <Select
                    value={editableEntity[key] as string}
                    onChange={(e) => {
                      setEditableEntity({ ...editableEntity, [key]: e.target.value });
                    }}
                  >
                    {Object.entries(selectChoices?.[key] || {}).map(([id, name]) => {
                      return (
                        <option key={id} value={id}>
                          {name}
                        </option>
                      );
                    })}
                  </Select>
                </Container>
              );
            } else return;
          })}
        </ModalBody>
        <ModalFooter>
          <Button mx={2} colorScheme="red" onClick={onClose}>
            Cancel
          </Button>
          <Button onClick={handleSave}>Save</Button>
        </ModalFooter>
      </ModalContent>
    </Modal>
  );
};

export default EntityModal;
