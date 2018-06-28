# Custom Dynamic DNS

I have a few machines which have a dynamic IPs. I like to be able to resolve these dynamic IPs via a domain I control, with the minimum amount of overhead. This creates a simple executable which uses the `name.com` API to update a DNS entry.

I call this via cron, every 5 minutes -

```*/5 * * * * /usr/local/bin/update-dns.sh```
