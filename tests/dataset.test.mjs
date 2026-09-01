import assert from "node:assert/strict";
import test from "node:test";
import { readFile } from "node:fs/promises";

const dataset=JSON.parse(await readFile(new URL("../data/schools.json",import.meta.url),"utf8"));
const geoExpected=new Map([["01KDVDNA003BF7FQZ9WWE8VPWE","Ashanti"],["01KDVDNA00AS06X8G4JFKNJHQD","Upper East"],["01KDVDNA00N6BFFK8VF5K8YXPW","Greater Accra"],["01KDVDNA00P1YSXFQ309TNM30Z","Upper West"],["01KDVDNA00V5W2KQKW61AH0TWY","Northern"],["01KDVDNA004Z43DWH8WRBWBMAG","Central"],["01KDVDNA0007GPXZWH5X216XXR","Central"],["01KDVDNA0073QFXJG0NQE8MPP6","Bono"],["01KDVDNA008YFMW192CVAT5ZS9","Eastern"],["01KDVDNA00C7MK5FG8WWKN2RC9","Greater Accra"],["01KDVDNA00BBYNG3K9WQS0Y02J","Volta"],["01KDVDNA001Y2TCAWTENGQT40P","Western"]]);

test("beta contains 16 unique provenance-complete public universities",()=>{assert.equal(dataset.schools.length,16);assert.equal(new Set(dataset.schools.map(s=>s.id)).size,16);for(const school of dataset.schools){assert.equal(school.type,"PUBLIC_UNIVERSITY");assert.equal(school.ownership,"PUBLIC");assert.deepEqual(school.sources,["gtec-current-institutions","ghanageo-2026.08.3-ulid"]);assert.ok(school.name&&school.locality&&school.region)}});
test("every pinned GhanaGeo reference resolves to the intended region snapshot",()=>{for(const school of dataset.schools)assert.equal(geoExpected.get(school.ghanaGeoPlaceId),school.region,`${school.id} geo reference`) });
test("search aliases never resolve to different records",()=>{const aliases=new Map;for(const school of dataset.schools)for(const raw of [school.id,school.name,...school.aliases]){const key=raw.toLowerCase();assert.ok(!aliases.has(key)||aliases.get(key)===school.id,`cross-record alias ${raw}`);aliases.set(key,school.id)}});
test("no sensitive or inferred campus fields are published",()=>{for(const school of dataset.schools)for(const forbidden of ["latitude","longitude","phone","email","fees","ranking","students","accreditationStatus"])assert.ok(!(forbidden in school),`${school.id} contains ${forbidden}`)});
