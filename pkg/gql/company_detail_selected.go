package gql

const CompanyDetailSelectedFields string = `
items {
  id
  name
  storeName
  descriptionEn
  descriptionId
  isSellerPkp
  createdAt
  updatedAt
  imageStorageId
  storageData {
    fileName
    fileType
    mime
    originalFilename
    path        
  }
  companyStores {
    id
    name
    slug
  }
  companyDocumentVerifications {
    status
  }
  institutions {
    code
  }
  province {
    id
    name
  }
  city {
    id
    name
  }
  district {
    id
    name
  }
  subDistrict {
    id
    name
  }
}

`
