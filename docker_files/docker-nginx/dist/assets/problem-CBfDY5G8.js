import{i as t,E as a}from"./request-Doo85ERT.js";const n=r=>t.get("/api2/problems",{params:r}).then(e=>e.data).catch(e=>(a({message:`服务错误: ${e}`,type:"error",plain:!0}),Promise.reject(e))),o=r=>t.post("/api2/problems",r).then(e=>!e.data).catch(e=>(a({message:`服务错误: ${e}`,type:"error",plain:!0}),Promise.reject(e))),p=r=>t.post("/api2/problems/update",r).then(e=>!e.data).catch(e=>(a({message:`服务错误: ${e}`,type:"error",plain:!0}),Promise.reject(e))),m=r=>t.post("/api2/problems/delete",r).then(e=>!e.data).catch(e=>(a({message:`服务错误: ${e}`,type:"error",plain:!0}),Promise.reject(e))),c=r=>t.get(`/api2/problems/${r}`).then(e=>e.data).catch(e=>(a({message:`服务错误: ${e}`,type:"error",plain:!0}),Promise.reject(e)));export{o as a,m as d,c as g,p as m,n as s};
