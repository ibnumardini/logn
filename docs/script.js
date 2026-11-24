(function () {
  'use strict';

  var sectionIds = [
    'overview',
    'installation',
    'configuration',
    'usage',
    'cron',
    'contributing'
  ];
  var activeSection = 'overview';
  var preserveInitialHash = false;
  var statusTimers = {};

  function updateActiveLinks(id) {
    activeSection = id;
    var links = document.querySelectorAll('#site-menu a[href^="#"]');
    links.forEach(function (link) {
      var target = link.getAttribute('href').slice(1);
      if (sectionIds.indexOf(target) !== -1) {
        if (target === activeSection) {
          link.setAttribute('aria-current', 'location');
        } else {
          link.removeAttribute('aria-current');
        }
      }
    });
  }

  function selectNearestSection() {
    var line = window.innerHeight * 0.35;
    var sections = sectionIds
      .map(function (id) { return document.getElementById(id); })
      .filter(Boolean);

    if (!sections.length) return;

    var atOrAbove = sections.filter(function (section) {
      return section.getBoundingClientRect().top <= line;
    });

    var candidate = atOrAbove.length
      ? atOrAbove.reduce(function (closest, section) {
          return section.getBoundingClientRect().top > closest.getBoundingClientRect().top
            ? section
            : closest;
        })
      : sections.find(function (section) {
          return section.getBoundingClientRect().top > line;
        });

    if (candidate && candidate.id !== activeSection) {
      updateActiveLinks(candidate.id);
    }
  }

  function setStatus(key, message, state) {
    var statusEl = document.getElementById('status-' + key);
    if (!statusEl) return;
    statusEl.textContent = message;
    statusEl.setAttribute('data-state', state || 'idle');
    clearTimeout(statusTimers[key]);
    if (state === 'copied' || state === 'error') {
      statusTimers[key] = setTimeout(function () {
        statusEl.textContent = '';
        statusEl.setAttribute('data-state', 'idle');
      }, 3000);
    }
  }

  function handleCopy(button) {
    var targetId = button.getAttribute('data-copy-target');
    var key = button.getAttribute('data-copy-key');
    var codeEl = targetId ? document.getElementById(targetId) : null;
    if (!codeEl || !key) return;

    var text = codeEl.innerText || codeEl.textContent || '';
    setStatus(key, '', 'idle');

    var fail = function () {
      setStatus(key, 'Clipboard unavailable. Select and copy the snippet manually.', 'error');
    };

    if (!navigator.clipboard || !window.isSecureContext) {
      fail();
      return;
    }

    try {
      var result = navigator.clipboard.writeText(text);
      Promise.resolve(result).then(
        function () {
          setStatus(key, 'Copied to clipboard.', 'copied');
        },
        fail
      );
    } catch (_err) {
      fail();
    }
  }

  function init() {
    var hash = window.location.hash.slice(1);
    if (sectionIds.indexOf(hash) !== -1) {
      activeSection = hash;
      preserveInitialHash = true;
    }
    updateActiveLinks(activeSection);

    window.addEventListener('hashchange', function () {
      var currentHash = window.location.hash.slice(1);
      if (sectionIds.indexOf(currentHash) !== -1) {
        updateActiveLinks(currentHash);
      }
    });

    if ('IntersectionObserver' in window) {
      var observer = new IntersectionObserver(
        function () {
          if (preserveInitialHash) {
            preserveInitialHash = false;
            return;
          }
          selectNearestSection();
        },
        { threshold: [0, 1] }
      );
      sectionIds.forEach(function (id) {
        var el = document.getElementById(id);
        if (el) observer.observe(el);
      });
    }

    window.addEventListener('scroll', function () {
      if (preserveInitialHash) {
        preserveInitialHash = false;
        return;
      }
      selectNearestSection();
    }, { passive: true });

    var copyButtons = document.querySelectorAll('button[data-copy-target]');
    copyButtons.forEach(function (button) {
      button.addEventListener('click', function () {
        handleCopy(button);
      });
    });
  }

  if (document.readyState === 'loading') {
    document.addEventListener('DOMContentLoaded', init);
  } else {
    init();
  }
})();
