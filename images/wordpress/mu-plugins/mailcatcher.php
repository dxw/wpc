<?php

add_action('phpmailer_init', function ($phpmailer) {
    $phpmailer->Host = 'mailcatcher';
    $phpmailer->Port = 1025;
    $phpmailer->SMTPAuth = false;
    $phpmailer->isSMTP();
});

add_filter('wp_mail_from', function ($mail) {
    return 'wordpress@localhost.invalid';
});
