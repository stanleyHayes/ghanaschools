export type School = Readonly<{ id:string; name:string; aliases:readonly string[]; type:string; ownership:string; locality:string; region:string; ghanaGeoPlaceId:string; locationNote?:string; sources:readonly string[] }>;
export type SchoolPage = Readonly<{ data:readonly School[]; count:number; datasetVersion:string; coverage:string }>;
export class GhanaSchoolsClient {
  constructor(private readonly baseUrl = "https://api-schools.digitalghana.dev") {}
  async search(filters: Readonly<{ q?:string; region?:string; type?:string }> = {}, signal?:AbortSignal): Promise<SchoolPage> { const url=new URL("/v1/schools",this.baseUrl); for(const [key,value] of Object.entries(filters)) if(value) url.searchParams.set(key,value); const response=await fetch(url,{signal}); if(!response.ok) throw new Error(`GhanaSchools request failed: ${response.status}`); return response.json() as Promise<SchoolPage>; }
  async get(id:string, signal?:AbortSignal): Promise<Readonly<{data:School;datasetVersion:string}>> { const response=await fetch(new URL(`/v1/schools/${encodeURIComponent(id)}`,this.baseUrl),{signal}); if(!response.ok) throw new Error(`GhanaSchools request failed: ${response.status}`); return response.json() as Promise<Readonly<{data:School;datasetVersion:string}>>; }
}
