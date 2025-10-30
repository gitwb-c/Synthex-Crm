<script lang="ts">
  import type { PageProps } from "./$types";
  import logo from "$lib/assets/Component 1 2.svg";
  import Input from "$lib/components/input.svelte";
  import { Mail } from "lucide-svelte";
  import Button from "$lib/components/button.svelte";
  import { notifications } from "$lib/stores/notifications";

  let { data, params }: PageProps = $props();

  const addNotification = async () => {
    const not = notifications;
    not.addNotification({
      message: Date.now().toLocaleString(),
      textColor: "white",
      backgroundColor: "red",
      timeToSkip: 2,
    });
  };
</script>

<main>
  <section id="company">
    <h1>Seja Bem vindo(a) a empresa</h1>
    <h1>{params.companyId}</h1>
    <img src={logo} alt="logo" />
  </section>
  <section id="forms">
    <form on:submit={addNotification}>
      <h1>Login</h1>
      <label>Email</label>
      <Input
        border="solid black 1px"
        leftIcon={Mail}
        type="email"
        width="100%"
      />

      <label>Senha</label>
      <Input
        isPassword={true}
        border="solid black 1px"
        type="password"
        width="100%"
      />
      <Button
        type="submit"
        fontSize=".9rem"
        fontWeight="600"
        title="Entrar"
        backgroundColor="var(--azul-marinho)"
        textColor="white"
      ></Button>
      <i>Esqueci minha senha</i>
    </form>
  </section>
</main>

<style>
  main {
    width: 100vw;
    height: 100vh;
    display: flex;

    & > section {
      width: 50%;
      height: 100%;
      display: flex;
      flex-direction: column;
    }
  }

  #company {
    text-align: center;
    padding: 2rem;
    gap: 2rem;
    background: linear-gradient(
      to bottom,
      var(--azul-claro),
      var(--azul-marinho)
    );

    & > h1 {
      font-size: 1.8rem;
      font-weight: 550;
      color: white;
    }

    & > img {
      max-height: 45%;
    }
  }

  form {
    width: 100%;
    height: 100%;
    display: flex;
    flex-direction: column;
    padding: 2.3rem;
    gap: 1rem;

    & > h1 {
      font-size: 2.2rem;
      font-weight: 500;
    }
  }

  i:hover {
    cursor: pointer;
    text-decoration-line: underline;
  }
</style>
