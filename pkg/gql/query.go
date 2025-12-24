package gql

const CompanyDetail = `
query ($query: QueryOption) {
	companyDetail(query: $query) {
    %%_SET_RESPONDS_FIELD%%
  }  
}
`
