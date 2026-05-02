<script lang="ts">
    import '../app.css';
    import './page.css';
    import { fade } from 'svelte/transition';
    import { onMount } from 'svelte';

    const DEFAULT_DOMAIN = 'b.4.0.c.7.0.4.1.a.2.ip6.arpa';
    let domain = DEFAULT_DOMAIN;
    let urlScheme = 'https:';

    onMount(() => {
        const p = window.location.protocol;
        urlScheme = p === 'http:' || p === 'https:' ? p : 'https:';
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
                        <div class="domain-example">{urlScheme}//{domain}</div>
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
