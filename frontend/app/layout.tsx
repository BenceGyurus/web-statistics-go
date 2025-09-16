import Head from "next/head";
import "./globals.css";

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="hu">
      <Head>
        <link rel="icon" href="/logo.png" sizes="any" />
      </Head>
      <body>
        {children}
      </body>
    </html>
  );
}
