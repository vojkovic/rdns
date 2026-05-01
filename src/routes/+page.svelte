<script lang="ts">
    import '../app.css';
    import { fade } from 'svelte/transition';
    import { onMount } from 'svelte';

    const DEFAULT_DOMAIN = 'b.4.0.c.7.0.4.1.a.2.ip6.arpa';
    let domain = DEFAULT_DOMAIN;

    onMount(() => {
        const h = window.location.hostname;
        if (h.endsWith('.ip6.arpa') || h.endsWith('.in-addr.arpa')) {
            domain = h;
        }
    });

    let currentSlide = 0;
    let showAAAA = false;
    const totalSlides = 4;

    function next() {
        if (currentSlide < totalSlides - 1) currentSlide++;
    }

    function prev() {
        if (currentSlide > 0) currentSlide--;
    }

    function handleKey(e: KeyboardEvent) {
        if (e.key === 'ArrowRight' || e.key === ' ') {
            e.preventDefault();
            next();
        } else if (e.key === 'ArrowLeft') {
            e.preventDefault();
            prev();
        }
    }
</script>

<svelte:window on:keydown={handleKey} />

<main>
    {#key currentSlide}
        <div class="slide" in:fade={{ duration: 200, delay: 100 }} out:fade={{ duration: 100 }}>

            {#if currentSlide === 0}
                <div class="slide-content center">
                    <h1 class="main-title">Did you know you can host a website on a reverse DNS domain?</h1>
                    <p class="subtitle">Click anywhere to find out how...</p>
                </div>

            {:else if currentSlide === 1}
                <div class="slide-content center">
                    <div class="card">
                        <div class="domain-example">https://{domain}</div>
                        <p>This is actually a valid domain name! <span class="emoji">🎉</span></p>
                    </div>
                </div>

            {:else if currentSlide === 2}
                <div class="slide-content center">
                    <div class="card text-left">
                        <h2>Wait, what? <span class="emoji">🤔</span></h2>
                        <p>Reverse DNS domains are usually just for PTR records, but you can actually add A records to them too (or even <button class="inline-toggle" on:click|stopPropagation={() => showAAAA = !showAAAA}>AAAA records</button>)!</p>
                        <div class="code-block">
                            <div>{domain}. IN A 23.140.20.94</div>
                            {#if showAAAA}
                                <div transition:fade={{ duration: 150 }}>{domain}. IN AAAA 2a14:7c0:4b00::94</div>
                            {/if}
                        </div>
                        <p>You could technically host your mail server on a reverse DNS domain!</p>
                        <p class="muted">Although, getting a <a href="https://vojk.au/posts/how_to_get_a_ip6_arpa_tls_certificate_part_2/" target="_blank" rel="noopener noreferrer" on:click|stopPropagation>TLS certificate</a> is tricky.</p>
                    </div>
                </div>

            {:else if currentSlide === 3}
                <div class="slide-content center">
                    <p class="credit-name"><a href="https://vojk.au" target="_blank" rel="noopener noreferrer" on:click|stopPropagation>Brock Vojkovic</a></p>
                    <p class="credit-line"><a href="https://as44354.net" target="_blank" rel="noopener noreferrer" on:click|stopPropagation>AS44354</a></p>
                </div>
            {/if}

        </div>
    {/key}

    <div class="nav">
        <button class="nav-btn" on:click|stopPropagation={prev} disabled={currentSlide === 0}>←</button>
        <div class="dots">
            {#each Array(totalSlides) as _, i}
                <button
                    class="dot"
                    class:active={i === currentSlide}
                    on:click|stopPropagation={() => currentSlide = i}
                    aria-label="Go to slide {i + 1}"
                ></button>
            {/each}
        </div>
        <button class="nav-btn" on:click|stopPropagation={next} disabled={currentSlide === totalSlides - 1}>→</button>
    </div>

    <!-- Click anywhere to advance -->
    <button class="click-area" on:click={next} aria-label="Next slide"></button>
</main>

<style>
    main {
        width: 100vw;
        height: 100vh;
        overflow: hidden;
        display: flex;
        align-items: center;
        justify-content: center;
        position: relative;
    }

    .click-area {
        position: absolute;
        inset: 0;
        z-index: 0;
        background: none;
        border: none;
        cursor: pointer;
        outline: none;
    }

    .slide {
        position: absolute;
        inset: 0;
        display: flex;
        align-items: center;
        justify-content: center;
        z-index: 1;
        pointer-events: none;
        padding: 2rem;
        box-sizing: border-box;
    }

    .slide-content {
        max-width: 640px;
        width: 100%;
        pointer-events: auto;
    }

    .slide-content.center {
        display: flex;
        flex-direction: column;
        align-items: center;
        text-align: center;
    }

    .main-title {
        font-size: 2.6rem;
        margin: 0 0 1.5rem 0;
        color: var(--text);
        line-height: 1.25;
        font-weight: 700;
    }

    .subtitle {
        font-size: 1.1rem;
        color: var(--text-muted);
        margin: 0;
    }

    .card {
        background: var(--surface);
        padding: 2rem;
        border-radius: 0.75rem;
        border: 1px solid var(--border);
        width: 100%;
        box-sizing: border-box;
    }

    .card.text-left {
        text-align: left;
    }

    .card h2 {
        color: var(--text);
        margin: 0 0 1rem 0;
        font-size: 1.5rem;
        font-weight: 600;
    }

    .card p {
        color: var(--text-muted);
        margin: 0 0 0.75rem 0;
        font-size: 1rem;
        line-height: 1.6;
    }

    .card p.muted {
        opacity: 0.7;
        font-size: 0.9rem;
    }

    .domain-example {
        font-family: 'Fira Code', monospace;
        font-size: 1.15rem;
        color: var(--accent);
        background: rgba(255, 255, 255, 0.04);
        padding: 0.75rem 1.25rem;
        border-radius: 0.5rem;
        border: 1px solid var(--border);
        margin-bottom: 1rem;
        letter-spacing: 0.02em;
    }

    .code-block {
        background: rgba(0, 0, 0, 0.5);
        padding: 1rem;
        border-radius: 0.5rem;
        font-family: 'Fira Code', monospace;
        font-size: 0.85rem;
        margin: 0.75rem 0;
        border: 1px solid var(--border);
        line-height: 1.8;
        color: var(--text);
    }

    .inline-toggle {
        background: none;
        border: none;
        color: var(--accent);
        cursor: pointer;
        font: inherit;
        padding: 0;
        text-decoration: underline;
        text-decoration-style: dotted;
        text-underline-offset: 2px;
    }

    .inline-toggle:hover {
        text-decoration-style: solid;
    }

    .emoji {
        font-size: 1.2rem;
    }

    a {
        color: var(--accent);
        text-decoration: underline;
        text-decoration-style: dotted;
        text-underline-offset: 2px;
    }

    a:hover {
        text-decoration-style: solid;
    }

    .attribution {
        margin-top: 0.5rem;
        font-size: 0.85rem;
        color: var(--text-muted);
    }

    .credit-line {
        font-size: 0.9rem;
        color: var(--text-muted);
        margin: 0 0 0.5rem 0;
    }

    .credit-name {
        font-size: 1.4rem;
        font-weight: 600;
        color: var(--text);
        margin: 0 0 0.25rem 0;
    }

    .credit-sub {
        font-size: 0.85rem;
        color: var(--text-muted);
        margin: 0;
    }

    .nav {
        position: fixed;
        bottom: 1.5rem;
        left: 50%;
        transform: translateX(-50%);
        display: flex;
        align-items: center;
        gap: 1rem;
        z-index: 10;
        pointer-events: auto;
    }

    .nav-btn {
        background: none;
        border: 1px solid var(--border);
        color: var(--text-muted);
        width: 32px;
        height: 32px;
        border-radius: 50%;
        cursor: pointer;
        font-size: 0.9rem;
        display: flex;
        align-items: center;
        justify-content: center;
        transition: color 0.2s, border-color 0.2s;
    }

    .nav-btn:hover:not(:disabled) {
        color: var(--text);
        border-color: var(--text-muted);
    }

    .nav-btn:disabled {
        opacity: 0.2;
        cursor: default;
    }

    .dots {
        display: flex;
        gap: 0.5rem;
    }

    .dot {
        width: 6px;
        height: 6px;
        border-radius: 50%;
        background: var(--border);
        border: none;
        padding: 0;
        cursor: pointer;
        transition: background 0.2s;
    }

    .dot.active {
        background: var(--accent);
    }

    @media (max-width: 768px) {
        .slide {
            padding: 1.5rem;
        }
        .main-title {
            font-size: 1.8rem;
        }
        .subtitle {
            font-size: 1rem;
        }
        .card {
            padding: 1.25rem;
        }
        .card h2 {
            font-size: 1.3rem;
        }
        .card p {
            font-size: 0.9rem;
        }
        .domain-example {
            font-size: 0.9rem;
            padding: 0.6rem 0.8rem;
        }
        .code-block {
            font-size: 0.7rem;
        }
        .nav {
            bottom: 1rem;
        }
    }

    @media (max-width: 480px) {
        .slide {
            padding: 1rem;
        }
        .main-title {
            font-size: 1.5rem;
        }
        .card {
            padding: 1rem;
        }
        .domain-example {
            font-size: 0.8rem;
        }
        .code-block {
            font-size: 0.6rem;
            overflow-x: auto;
        }
    }
</style>
