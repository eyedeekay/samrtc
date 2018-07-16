# samrtc, or, How I stopped worrying and am making WebTorrent work over i2p without asking permission.

**HIGHLY EXPERIMENTAL. IF I MAKE A MISTAKE SECURING THIS APPLICATION, I COULD**
**ACCIDENTALLY EXPOSE YOUR SAM PORT TO THE WHOLE I2P NETWORK. IF SOMEONE FINDS**
**IT, THEY CAN USE YOUR I2P ROUTER TO CREATE NEW CONNECTIONS, SEND MESSAGES,**
**AND A WHOLE BUNCH OF OTHER STUFF *AS YOU*. DO NOT DEPEND ON IT FOR ANONYMITY**
**UNTIL DESTINATION WHITELISTING IS IMPLEMENTED COMPLETELY, WITH 100% TEST**
**COVERAGE AND LOTS OF DOGFOODING BY ME.** That said, I can't stop you.

Enable webRTC on i2p. Create a tunnel to your own SAM port over i2p with
whitelisting for the local destination. DANGER: Of all the potentially stupid
things I have done with i2p, this is probably the most risky. It might also be
unnecessary, but if I'm right about the security of most browsers people attach
to i2p, then it might be a good idea.

First of all, sure, it's possible to proxy WebRTC over a regular SOCKS proxy or
whatever. But when it comes to the more interesting WebRTC applications, it
appears to me that this approach would be very limiting. Many WebTorrent-based
applications would be problematic in this scenario, for instance. A better idea,
I think, would be to make it possible for WebRTC Applications to use the safer
i2p API's like the SAM port to make connections. But the obvious problem
arises, when in a properly configured browser, unproxied connections to the
localhost are disabled, which means that javascript in the browser can't talk
to the SAM port directly, and that's the right thing to do. So, how do we
make it possible for our browser, and **only** our browser, to talk to our SAM
port, and **only** our SAM port? Using the SAM bridge to forward the
connection, over i2p, using a 0-hop tunnel for the SAM port service with
destination whitelisting for the local http proxy. That way, your SAM bridge is
just a really fast service you can access via a .b32 address, which will
disallow any connection that doesn't originate from your i2p http proxy. But I
think I can still do better. In my [destination-isolation](https://github.com/eyedeekay/si-i2p-plugin)
project, I establish browser connections over the SAM bridge as well. It could
be used to allow the client connection to the SAM port over i2p to use zero
hops. This will allow you to access your local SAM port, via an i2p destination,
that only corresponds to the local port intended to access the SAM port.

The drawback, however, is it's only half of what you need. The other half is,
unfortunately, that the procedure is no longer generic. Instead, you need to
modify the webRTC applications and libraries to use the SAM Bridge to establish
connections over i2p instead.

