<script lang="ts">
    import '../app.css';
    import IPv6Animation from '$lib/components/IPv6Animation.svelte';

    let showAAAA = false;
    let codeBlock: HTMLElement;

    function toggleAAAA() {
        showAAAA = !showAAAA;
        if (codeBlock) {
            codeBlock.style.height = showAAAA ? 'auto' : '0';
            codeBlock.style.opacity = showAAAA ? '1' : '0';
        }
    }
</script>

<main>
    <div class="container">
        <h1 class="main-title">Did you know you can host a website on a reverse DNS domain?</h1>
        <p class="subtitle">Read on to find out how!</p>

        <div class="fun-fact">
            <div class="domain-example">b.4.0.c.7.0.4.1.a.2.ip6.arpa</div>
            <p>This is actually a valid domain name! <span class="emoji">🎉</span></p>
        </div>

        <div class="explanation">
            <h2>Wait, what? <span class="emoji">🤔</span></h2>
            <p>While reverse DNS domains are typically used for PTR records (you know, the boring stuff), there's nothing stopping you from adding A records to them as well. <a href="#" on:click|preventDefault={toggleAAAA} class="aaaa-link">(or AAAA)</a></p>
            
            <div class="code-block">
                b.4.0.c.7.0.4.1.a.2.ip6.arpa. IN A 67.219.98.26
                <div class="aaaa-record" bind:this={codeBlock}>
                    b.4.0.c.7.0.4.1.a.2.ip6.arpa. IN AAAA 2a14:7c0:4b10::1
                </div>
            </div>

            <p>That's right - you could technically host your personal blog on a reverse DNS domain if you wanted to!</p>
            <p>Although, getting a TLS certificate for a reverse DNS domain is a bit tricky. Let's Encrypt won't issue a certificate. <span class="emoji">🤷‍♂️</span></p>
        </div>

        <IPv6Animation />
    </div>
</main>

<style>
    :root {
        --background: hsl(0 0% 6%);
        --primary: hsl(255, 78%, 55%);
        --secondary: hsl(255, 78%, 50%);
        --accent: hsl(0 0% 10%);
        --text-color: #ffffff;
        --subtext-color: #eee;
        --hover-color: #aaa;
        --button-text-color: #ffffff;
        --table-border: #888;
        --table-header-bg: #171717;
    }

    body {
        font-family: 'Inter', -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif;
        line-height: 1.6;
        margin: 0;
        padding: 0;
        background-color: var(--background);
        color: var(--text-color);
        min-height: 100vh;
        display: flex;
        flex-direction: column;
        align-items: center;
        justify-content: center;
        text-align: center;
        background-image: 
            radial-gradient(circle at 25% 25%, rgba(255, 255, 255, 0.05) 0%, transparent 50%),
            radial-gradient(circle at 75% 75%, rgba(255, 255, 255, 0.05) 0%, transparent 50%);
        background-attachment: fixed;
    }

    .container {
        max-width: 800px;
        padding: 2rem;
    }

    .main-title {
        font-size: 3.5rem;
        margin-bottom: 1rem;
        background: linear-gradient(135deg, var(--primary), var(--secondary));
        -webkit-background-clip: text;
        -webkit-text-fill-color: transparent;
        line-height: 1.2;
    }

    .subtitle {
        font-size: 1.5rem;
        color: var(--subtext-color);
        margin-bottom: 3rem;
        font-style: italic;
    }

    .fun-fact {
        background: var(--accent);
        padding: 2rem;
        border-radius: 1rem;
        margin-bottom: 2rem;
        border: 1px solid var(--table-border);
        box-shadow: 0 4px 6px -1px rgba(0, 0, 0, 0.3);
        transform: rotate(-1deg);
    }

    .domain-example {
        font-family: 'Fira Code', monospace;
        font-size: 1.2rem;
        color: var(--primary);
        margin: 1rem 0;
        word-break: break-all;
        background: var(--background);
        padding: 1rem;
        border-radius: 0.5rem;
        border: 1px solid var(--table-border);
    }

    .explanation {
        background: var(--background);
        padding: 2rem;
        border-radius: 1rem;
        margin: 2rem 0;
        border: 1px solid var(--table-border);
        text-align: left;
        transform: rotate(1deg);
    }

    .explanation h2 {
        color: var(--primary);
        margin-top: 0;
        font-size: 1.8rem;
    }

    .explanation p {
        color: var(--subtext-color);
        margin-bottom: 1rem;
        font-size: 1.1rem;
    }

    .code-block {
        background: var(--table-header-bg);
        padding: 1rem;
        border-radius: 0.5rem;
        font-family: 'Fira Code', monospace;
        margin: 1rem 0;
        border: 1px solid var(--table-border);
        transform: rotate(-0.5deg);
    }

    .aaaa-record {
        height: 0;
        opacity: 0;
        overflow: hidden;
        transition: all 0.3s ease-in-out;
    }

    .aaaa-link {
        color: var(--primary);
        text-decoration: none;
        border-bottom: 1px dashed var(--primary);
        cursor: pointer;
    }

    .aaaa-link:hover {
        color: var(--secondary);
        border-bottom-color: var(--secondary);
    }

    .emoji {
        font-size: 1.5rem;
        margin: 0 0.2rem;
    }

    @media (max-width: 768px) {
        .container {
            padding: 1rem;
        }
        
        .main-title {
            font-size: 2.5rem;
        }

        .fun-fact, .explanation {
            transform: none;
        }
    }
</style>
