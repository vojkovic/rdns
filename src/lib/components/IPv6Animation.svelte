<script lang="ts">
    import { onMount } from 'svelte';

    let currentStep = 0;
    let previousStep = 0;
    let isPlaying = false;
    let animationInterval: number;
    let container: HTMLElement;
    let isVisible = false;
    let containerWidth = 0;
    let fontSize = 1.2;

    const steps = [
        {
            title: "Original IPv6 Prefix",
            content: "2a14:07c0:4b00::/40",
            description: "This is the delegated IPv6 prefix with a /40 mask",
            highlight: [0, 1, 2, 3, 5, 6, 7, 8, 10, 11]
        },
        {
            title: "Split into Nibbles",
            content: "2 a 1 4 0 7 c 0 4 b 0 0",
            description: "Each hexadecimal nibble is separated for processing",
            highlight: [0, 2, 4, 6, 8, 10, 12, 14, 16, 18]
        },
        {
            title: "Trim to Prefix Length",
            content: "2 a 1 4 0 7 c 0 4 b",
            description: "Nibbles outside the prefix length are removed",
            highlight: [0, 2, 4, 6, 8, 10, 12, 14, 16, 18]
        },
        {
            title: "Reverse Order",
            content: "b 4 0 c 7 0 4 1 a 2",
            description: "The nibbles are reversed for DNS lookup order",
            highlight: [0, 2, 4, 6, 8, 10, 12, 14, 16, 18]
        },
        {
            title: "Add Dots",
            content: "b.4.0.c.7.0.4.1.a.2",
            description: "Dots are added between each nibble for DNS format",
            highlight: [0, 2, 4, 6, 8, 10, 12, 14, 16, 18]
        },
        {
            title: "Add .ip6.arpa",
            content: "b.4.0.c.7.0.4.1.a.2.ip6.arpa",
            description: "The .ip6.arpa suffix is added to complete the reverse DNS domain",
            highlight: [0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 21, 22, 24, 25, 26, 27]
        },
        {
            title: "Done",
            content: "b.4.0.c.7.0.4.1.a.2.ip6.arpa",
            description: "Your reverse DNS domain is ready!",
            highlight: [0]
        }
    ];

    function updateFontSize() {
        if (!container) return;
        
        const display = container.querySelector('.ipv6-display');
        if (!display) return;

        const content = steps[currentStep].content;
        const charCount = content.length;
        const availableWidth = containerWidth - 40;
        const minCharWidth = 1.2;
        const maxFontSize = 1.2;
        const minFontSize = 0.5;

        if (currentStep === steps.length - 2) {
            fontSize = Math.min(0.8, maxFontSize);
            return;
        }

        const optimalFontSize = Math.min(
            maxFontSize,
            Math.max(minFontSize, availableWidth / (charCount * minCharWidth))
        );

        fontSize = optimalFontSize;
    }

    function handleResize() {
        if (container) {
            containerWidth = container.offsetWidth;
            updateFontSize();
        }
    }

    function resetAnimationState() {
        const display = container?.querySelector('.ipv6-display');
        if (display) {
            display.style.animation = 'none';
            display.style.transform = '';
            display.style.opacity = '';
            
            display.offsetHeight;
            
            const characters = display.querySelectorAll('.character');
            characters.forEach(char => {
                char.style.animation = 'none';
                char.style.transform = '';
                char.style.opacity = '';
                char.offsetHeight;
            });
        }
    }

    function playAnimation() {
        if (!isPlaying) {
            isPlaying = true;
            animationInterval = setInterval(() => {
                resetAnimationState();
                previousStep = currentStep;
                currentStep = (currentStep + 1) % steps.length;
                updateFontSize();
            }, 3000);
        }
    }

    function pauseAnimation() {
        isPlaying = false;
        clearInterval(animationInterval);
    }

    function handleIntersection(entries: IntersectionObserverEntry[]) {
        const entry = entries[0];
        isVisible = entry.isIntersecting;
        if (isVisible) {
            resetAnimationState();
            playAnimation();
        } else {
            pauseAnimation();
        }
    }

    onMount(() => {
        const observer = new IntersectionObserver(handleIntersection, {
            threshold: 0.5
        });
        
        if (container) {
            observer.observe(container);
            handleResize();
        }

        window.addEventListener('resize', handleResize);

        return () => {
            observer.disconnect();
            clearInterval(animationInterval);
            window.removeEventListener('resize', handleResize);
        };
    });
</script>

<div class="animation-container" bind:this={container}>
    <h2 class="step-title">{steps[currentStep].title}</h2>

    <div class="step-container">
        <div class="step-content">
            <div class="animation-stage">
                {#if currentStep === steps.length - 1}
                    <div class="final-domain">
                        {steps[currentStep].content}
                    </div>
                {:else}
                    <div class="ipv6-display" style="font-size: {fontSize}rem" key={currentStep}>
                        {#each steps[currentStep].content.split('') as char, i}
                            <span 
                                class="character" 
                                class:digit={char !== ':' && char !== '.' && char !== ' ' && char !== '/'}
                                class:separator={char === ':' || char === '.'}
                                class:space={char === ' '}
                                class:highlight={steps[currentStep].highlight.includes(i)}
                                style="--delay: {i * 50}ms"
                            >
                                {char}
                            </span>
                        {/each}
                    </div>
                {/if}
            </div>
        </div>
    </div>

    <p class="description">{steps[currentStep].description}</p>

    <div class="progress-bar">
        <div class="progress" style="width: {((currentStep + 1) / steps.length) * 100}%"></div>
    </div>
</div>

<style>
    .animation-container {
        background: var(--accent);
        padding: 2rem;
        border-radius: 1.5rem;
        margin: 2rem 0;
        border: 1px solid var(--table-border);
        box-shadow: 0 8px 32px rgba(0, 0, 0, 0.1);
        transform: rotate(0.5deg);
        display: flex;
        flex-direction: column;
        gap: 1.5rem;
    }

    .step-title {
        color: var(--text-color);
        margin: 0;
        font-size: 1.8rem;
        text-align: center;
    }

    .progress-bar {
        width: 100%;
        height: 6px;
        background: rgba(255, 255, 255, 0.1);
        border-radius: 3px;
        overflow: hidden;
        margin-top: auto;
    }

    .progress {
        height: 100%;
        background: var(--primary);
        transition: width 0.5s cubic-bezier(0.4, 0, 0.2, 1);
    }

    .step-container {
        min-height: 120px;
    }

    .animation-stage {
        width: 100%;
        display: flex;
        justify-content: center;
        align-items: center;
        padding: 1.5rem 0;
        overflow-x: auto;
        overflow-y: hidden;
    }

    .ipv6-display {
        display: flex;
        flex-wrap: nowrap;
        justify-content: center;
        align-items: center;
        gap: 0.15rem;
        font-family: 'Fira Code', monospace;
        line-height: 1;
        white-space: nowrap;
        padding: 0 1rem;
        transition: font-size 0.3s ease;
    }

    .character {
        display: inline-flex;
        align-items: center;
        justify-content: center;
        flex-shrink: 0;
        width: 1.2em;
        height: 1.2em;
        text-align: center;
        position: relative;
        will-change: transform, opacity;
    }

    .character.digit {
        background: rgba(255, 255, 255, 0.1);
        color: var(--text-color);
        padding: 0.2rem 0.3rem;
        border-radius: 0.4rem;
        border: 1px solid rgba(255, 255, 255, 0.1);
    }

    .character.separator {
        color: var(--text-color);
        font-weight: 500;
    }

    .character.space {
        width: 0.5em;
    }

    .character.highlight {
        background: var(--primary);
        color: var(--button-text-color);
        border-color: var(--primary);
    }
    
    .character.highlight.full-domain {
        padding: 0.5rem 1rem;
        border-radius: 0.5rem;
        width: auto;
        height: auto;
    }

    .description {
        color: var(--subtext-color);
        text-align: center;
        margin: 0;
        font-size: 1.1rem;
    }

    .final-domain {
        background: var(--primary);
        color: var(--button-text-color);
        padding: 0.8rem 1.5rem;
        border-radius: 0.5rem;
        font-family: 'Fira Code', monospace;
        font-size: 1.4rem;
        white-space: nowrap;
        box-shadow: 0 4px 8px rgba(0, 0, 0, 0.2);
        animation: fadeIn 0.5s ease-out;
    }

    @keyframes fadeIn {
        from { opacity: 0; transform: translateY(10px); }
        to { opacity: 1; transform: translateY(0); }
    }

    @media (max-width: 768px) {
        .animation-container {
            padding: 1rem;
            transform: none;
        }

        .ipv6-display {
            font-size: 0.8rem;
        }
    }
</style> 