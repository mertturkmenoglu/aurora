Aurora E-commerce
---
# What is it?
* Aurora is a test/learning project.
* We (my other teammates don't commit, they are not working. Not sure why.) wanted
to try different technologies. So we created this dummy project.
* It's not complete, it never will be complete. I'm just having fun.
* Go, Nuxt.js, Postgres, and Docker are the fundamental technologies.
# Requirements
* Go 1.20 or later
* Docker
* Node.js 18 or later
* Yarn
* Doppler CLI
---
# Installation
* Clone the repository.
* Configure Doppler CLI using `doppler.yaml` inside the `apps/api` folder.
* Run `yarn` inside `apps/web` folder.
---
# Running
* Run `docker-compose up` inside the root folder.
* Run `CREATE EXTENSION IF NOT EXISTS "uuid-ossp";` for your database to setup `uuid_generate_v4` extension.
* Run `yarn dev` inside `apps/web` folder.
* Run `doppler run -- air` inside `apps/api` folder.
# FAQ (Srsly)
* Q: What branching strategy are you using?  
  A: It's called Imdb-branching. I'm creating random branch names from (*mostly*) movies. 
* Q: Is this a joke to you?  
  A: Yes?
* Q: Why are you having fun? Stop having fun and code like an adult. Clean code, SOLID, clean architecture, internals folder, cmd folder, blah blah.  
  A: [Codes happily while listening to Baldur's Gate 3 OST] What mate? You said something? I didn't catch it, I was listening to music.
* Q: How do you run this thing? What should be the environment variables?  
  A: That's the neat part: you don't. Idk, I don't remember. I copy pastad to Doppler and forget about it.
