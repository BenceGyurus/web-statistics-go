"use client"

import Selector from "@/components/selector/Selector.component";
import { useEffect, useState } from "react";

export default function Home() {

  const [sites, setSites] = useState([]);

  useEffect(() => {
    fetch('http://localhost:3001/api/v1/get-sites', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
    })
      .then(response => response.json())
      .then(data => setSites(data.sites || []))
      .catch(error => console.error('Error fetching stats:', error));
  }, []);

  return (
    <div>
      <div>
          <h1>Statisztikák</h1>
          <div>
              <Selector options={sites.filter(site => {return site!==""})} selected={""} onSelect={(value) => console.log(value)} />
          </div>
      </div>
    </div>
  );
}
