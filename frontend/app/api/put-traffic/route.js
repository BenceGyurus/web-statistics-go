// pages/api/proxy.js
export default async function handler(req, res) {
    res.setHeader('Access-Control-Allow-Origin', '*');
    res.setHeader('Access-Control-Allow-Methods', 'GET,POST,OPTIONS');
    res.setHeader('Access-Control-Allow-Headers', 'Content-Type, Authorization');
  
    if (req.method === 'OPTIONS') {
      res.status(200).end();
      return;
    }

    const url = process.ENV.BACKEND + "/" + req.url.replace('/api/', '/api/v1');
    const backendRes = await fetch(url, {
      method: req.method,
      headers: { ...req.headers },
      body: req.method !== 'GET' && req.method !== 'HEAD' ? req.body : undefined,
    });
    const data = await backendRes.text();
  
    res.status(backendRes.status).send(data);
  }
  