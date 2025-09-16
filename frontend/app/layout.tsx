import "./globals.css";

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="hu">
      <link rel="icon" href="/logo.png" sizes="any" />
      <body>
        {children}
      </body>
    </html>
  );
}
