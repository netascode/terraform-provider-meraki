data "meraki_switch_stack_routing_interface_dhcp" "example" {
  network_id      = "L_123456"
  switch_stack_id = "1234"
  interface_id    = "1234"
}
