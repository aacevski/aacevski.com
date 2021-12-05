export default async function Fetcher(...args: Parameters<typeof fetch>) {
  const res = await fetch(...args);
  return res.json();
}
