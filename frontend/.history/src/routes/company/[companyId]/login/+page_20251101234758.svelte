<script lang="ts">
  import type { PageProps } from "./$types";
  import logo from "$lib/assets/Component 1 2.svg";
  import Input from "$lib/components/input.svelte";
  import { Mail, Building } from "lucide-svelte";
  import Button from "$lib/components/button.svelte";
  import { notifications } from "$lib/stores/notifications";
  import { goto } from "$app/navigation";
  import Helper from "$lib/helper";
  import { BACKEND_API_URL } from "$env/static/private";

  let { data, params }: PageProps = $props();

  let loadingSubmit = $state(false);

  const delay = (seconds: number) =>
    new Promise((resolve) => setTimeout(resolve, seconds * 1000));

  const submit = async () => {
    if (loadingSubmit) return;

    loadingSubmit = true;

    if (email == "" || password == "") {
      notifications.addNotification({
        message: "Campos faltando!",
        textColor: "white",
        backgroundColor: "var(--vermelho)",
        timeToSkip: 2,
      });
      loadingSubmit = false;
      return;
    }

    const res = await Helper.fetchUrl(
      `${BACKEND_API_URL}/auth/login/${params.companyId}`,
      { method: "POST", body: JSON.stringify({ email, password }) }
    );

    goto(`/company/${params.companyId}/pipeline/Qualificação`);
    if (res.statusCode == 200) loadingSubmit = false;
  };

  let email = $state(""),
    password = $state("");
</script>

<main>
  <section id="company">
    <h1>Synthex CRM</h1>
    <!-- <div>
      <Building size="5rem"></Building>
      <h2>
        {params.companyId}
      </h2>
    </div> -->
  </section>
  <section id="forms">
    <!-- svelte-ignore event_directive_deprecated -->
    <form on:submit={submit}>
      <h1>Login</h1>
      <!-- svelte-ignore a11y_label_has_associated_control -->
      <label>Email</label>
      <Input
        border="solid black 1px"
        leftIcon={Mail}
        type="email"
        width="100%"
        iconWidth="6%"
        inputWidth="94%"
        bind:value={email}
      />

      <!-- svelte-ignore a11y_label_has_associated_control -->
      <label>Senha</label>
      <Input
        isPassword={true}
        border="solid black 1px"
        type="password"
        width="100%"
        iconWidth="6%"
        inputWidth="94%"
        bind:value={password}
      />
      <Button
        type="submit"
        fontSize=".9rem"
        fontWeight="600"
        title={loadingSubmit ? "Carregando..." : "Entrar"}
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
    background-image: url(https://images.unsplash.com/photo-1531512073830-ba890ca4eba2?ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxzZWFyY2h8Nnx8cGxhbm8lMjBkZSUyMGZ1bmRvJTIwcGFpc2FnZW18ZW58MHx8MHx8fDA%3D&fm=jpg&q=60&w=3000);
    background-repeat: no-repeat;
    background-size: cover;
    justify-content: space-between;
  }

  #company {
    width: 50%;
    height: 100%;
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    gap: 2.9rem;

    & > h1 {
      color: white;
      letter-spacing: 0.2rem;
      font-weight: 600;
      background-color: #1c5cc2dc;
      padding: 1rem;
      box-shadow: 0 0 2rem rgba(0, 0, 0, 0.6);
    }

    & > div {
      display: flex;
      flex-direction: column;
      align-items: center;
      background-color: var(--azul-palido);
      padding: 1rem;
      color: var(--azul-marinho);
      gap: 1rem;
      box-shadow: 0 0 2rem rgba(0, 0, 0, 0.6);
      border-radius: 1rem;
    }
  }

  #forms {
    width: 50%;
    height: 100vh;
    display: flex;
    justify-content: center;
    align-items: center;
    flex-direction: column;
    background-color: transparent;
  }

  form {
    background-color: white;
    box-shadow: 0 0 0.4rem rgba(0, 0, 0, 0.6);
    width: 60%;
    height: 60%;
    display: flex;
    flex-direction: column;
    padding: 2rem;
    gap: 0.9rem;
    font-size: 0.9rem;

    & > h1 {
      font-size: 2rem;
      font-weight: 500;
    }

    @media (min-width: 1400px) {
      width: 55%;
      height: 50%;

      font-size: 1rem;

      & > h1 {
        font-size: 2.2rem;
      }
    }
  }

  i:hover {
    cursor: pointer;
    text-decoration-line: underline;
  }
</style>
